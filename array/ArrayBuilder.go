package main
////////////////////////
import "fmt"
import "time"
const ramLength = 5000


///////////////////////////////////////

type ArrayIntInterface interface {
  push(int)
}

type AInt struct {
    ramLength *int
    rStart int
    rPointer int
    rEnd int
    step int
    len int
    pointer int
    follow *int
    ram *[ramLength]int
}

func (this *AInt) push(v int)  {
    if this.rPointer  >= this.rEnd {
         length :=  this.len + this.step
         if *this.follow != this.rEnd {
            b := *this.follow + length
            this.rPointer = *this.follow + this.pointer
            if b > ramLength {
                *this.follow = 0
            }
            end := b    
            copy(this.ram[*this.follow:],this.ram[this.rStart:this.rEnd]) 
            this.rStart = *this.follow 
            this.rEnd = end
            *this.follow = end
         }else {
            b := this.rEnd + length
            if ( b > ramLength) {
               *this.follow = 0
                copy(this.ram[*this.follow:],this.ram[this.rStart:this.rEnd])
                this.rStart = 0
                this.rEnd = length
                this.rPointer = *this.follow + this.pointer
            }else{
               this.rEnd =b
               *this.follow = this.rEnd
            }
         }
    }
    this.ram[this.rPointer] = v
    this.rPointer++
    this.pointer ++     
}

///////////////////////////////////////

type ArrayStringInterface interface {
    push(string) 
    join(string) string
}

type AString struct {
    rStart int
    rPointer int
    rEnd int
    step int
    len int
    strlen int
    pointer int
    follow *int
    ramLength *int
    ram *[ramLength]string
}


func (this *AString) push(v string)  {
    if this.rPointer  >= this.rEnd {
         length :=  this.len + this.step
         if *this.follow != this.rEnd {
            b := *this.follow + length
            this.rPointer = *this.follow + this.pointer
            if b > ramLength {
                *this.follow = 0
            }
          
            end := b    
            copy(this.ram[*this.follow:],this.ram[this.rStart:this.rEnd]) 
            this.rStart = *this.follow 
            this.rEnd = end
            *this.follow = end
         }else {
            b := this.rEnd + length
            if (b > ramLength) {
               *this.follow = 0
                copy(this.ram[*this.follow:],this.ram[this.rStart:this.rEnd])
                this.rStart = 0
                this.rEnd = length
                this.rPointer = *this.follow + this.pointer
            }else{
               this.rEnd =b
               *this.follow = this.rEnd
            }
         }
    }
    this.strlen += len(v)
    this.ram[this.rPointer] = v
    this.rPointer++
    this.pointer ++     
}

func (this *AString) join(a string) string {   

     /// 247127700
     
     t:= make([]byte, this.strlen + (this.len*len(a)))
     j:=0
     j +=  copy(t[j:],this.ram[this.rStart])
     for i:=this.rStart+1 ; i < this.rPointer ; i++ {
              j +=  copy(t[j:],a)
              j +=  copy(t[j:],this.ram[i])
     
     }
     return  string(t[:j])
}

///////////////////////////////////////


type ArrayBuilderInterface interface {
    string(int) ArrayStringInterface
    int(int) ArrayIntInterface
}

type ArrayBuild struct {
    ramString [ramLength]string
    ramInt [ramLength]int
    sString int
    sInt int
}


func (this *ArrayBuild) int(length int) ArrayIntInterface {
    b := this.sInt + length 
    if b > ramLength {
        this.sInt = 0
    }
    a := new(AInt)
    a.step = length
    a.len = length
    a.rStart   = this.sInt
    a.rPointer = this.sInt
    a.pointer = 0 
    a.rEnd   = b
    a.follow = &this.sInt
    a.ram = &this.ramInt
    this.sInt = b
    return (a)
}

func (this *ArrayBuild) string(length int) ArrayStringInterface {
    b := this.sString + length 
    if b > ramLength {
        this.sString = 0
    }
    a := new(AString)
    a.step = length
    a.len = length
    a.rStart   = this.sString
    a.rPointer = this.sString
    a.pointer = 0 
    a.rEnd   = b
    a.strlen = 0
    a.follow = &this.sString
    a.ram = &this.ramString
    this.sString = b
    return (a)
}


func ArrayBuilder(length int) ArrayBuilderInterface{
        a:=  new(ArrayBuild)
        a.ramString  = [ramLength]string{}
        a.ramInt  = [ramLength]int{}
        a.sString = 0
        a.sInt = 0
        return (a)
}

///////////////////////////////////////

func main() {

    a:= ArrayBuilder(700)

    b:= a.string(2)

    b.push("hello world")
    b.push("hello world")
    b.push("go go")
    fmt.Println(b.join(" , ")) 

    var t1 int64
   limit := 2000000
   // var c1 int64 = 0
   // var c2 int64 = 0
   var c3 int64 = 0
    var i int


 
    /// 155989100
    /// 149714600
    i = 0
 
    a.string(500)
    for {
            if i >= limit { break }  
            i ++
            t1 = time.Now().UnixNano()
            r:= a.string(25)
            r.push("c++ sqdsqdsq azdzd azda zad zazdz adaz zd az dzd ")
        
            c3 += time.Now().UnixNano()-t1
             
    } 
    fmt.Println(c3) 

    

}