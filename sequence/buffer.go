package main
import (
    "time"
    "fmt"
)

////////////////////////////

type Part struct {
    uid int
    pEnd int
    lock bool 
    tampon []byte
    buffer *Buffer
    length int
    step int
}

type PartInterface interface {
    push(string) 
    resolve()
}

func (this *Part) push (a string) {
    length := this.pEnd + len(a)
    if  length > this.length {
         extend := length + this.step
         t:= make([]byte, extend)
         copy(t[0:],this.tampon[0:]);
         this.tampon = t
         this.length = extend
    }
    if !this.lock {
        this.pEnd += copy(this.tampon[this.pEnd:],a);
    }
 
    
}
func (this *Part) resolve () {
    this.buffer.resolve(this.uid)
    this.lock = true
}

func (this *Part) getLength() int{
    return this.pEnd
}


func (this *Part) getTampon() []byte{
    a := make([]byte,this.pEnd)
    copy(a[0:],this.tampon[0:this.pEnd])
    return a
}


type Buffer struct {
    parts   []*Part
    uid int
    nbPart  int
    nbPartDone  int
    promise func(string)
    
}

type BufferInterface interface {
    part(int) PartInterface
    resolve(int)
    
}


func (this *Buffer) part(length int) PartInterface{
    if this.uid > this.nbPart {
         fmt.Println("end of distribution")
    } 
    a := new(Part)
    a.uid    = this.uid
    a.length = length
    a.step   = length
    a.lock   = false
    a.tampon = make([]byte,length)
    a.buffer = this
    this.parts[this.uid] = a
    this.uid ++
    
    return a
}

func (this *Buffer) resolve(a int){
    this.nbPartDone ++ 
    if this.nbPartDone >= this.nbPart {
          l := 0
          for i:=0; i<this.nbPart ; i++ {
             l +=this.parts[i].getLength()
          } 
          r := make([]byte,l)
          w := 0
          
          for i:=0; i<this.nbPart ; i++ {
             w += copy(r[w:],this.parts[i].getTampon())
          }
          this.promise(string(r))
    }
}

func buffer(nbPart int,promise func(string)) BufferInterface{
    a := new(Buffer)
    a.parts = make([]*Part,nbPart)
    a.promise = promise
    a.nbPart = nbPart
    return (a)
}

////////////////////////////////

func thread2(a PartInterface){
  
    a.push("<ul>")
    for i:=0;  i<5 ; i++ {
        a.push("<li>c</li>")
    }
    a.push("</ul>")
    a.resolve()
}

func thread(a PartInterface){
    time.Sleep(time.Second * 1)
    a.push("<h1>Title</h1>")
    a.resolve()
}

func main() {

    a:= buffer( 2 , func(r string){
            fmt.Println(r)
    })
  
    go  thread  (a.part(50))
    go  thread2 (a.part(20))
    
    time.Sleep(time.Second * 20)
}