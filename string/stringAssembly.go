package main
////////////////////////
import "fmt"

const ramLength = 200
const ramBloc = 2
const space = 4
const nbBloc = ramLength/ramBloc

////////////////////////

type partInterface interface {
    push(string) 
   
    log() 
}

type AString struct {
    rStart int
    rEnd int
    rBloc int
    rPointer int
    cuid int
  
    len int
    strlen int
    pointer int

    ram *[ramLength]byte
    extend func()
}

func (this *AString) log()  {
        fmt.Println("**********")
        fmt.Println(this.cuid)
        fmt.Println("---POS---")
        fmt.Println(this.rStart)
        fmt.Println(this.rEnd)
        fmt.Println("------")
        fmt.Println(this.len)
        fmt.Println("---POINT---")
        fmt.Println(this.rPointer)
        fmt.Println("**********")
}

func (this *AString) push(v string)  {
        if this.rPointer + len(v)  > this.rEnd {
            this.extend()
        }
        this.strlen += len(v)
        copy(this.ram[this.rPointer:],v)
       
        this.rPointer+=len(v)
       
        this.pointer +=len(v)   
        //this.log()  
}



///////////////////////////////////////


type BufferInterface interface {
    part() partInterface

}

type Buffer struct {
    ramString [ramLength]byte
    sString int
    sInt int
    cmp int
    cuid int
    length int
    partString [ramLength/ramBloc]int
    
    sPartString int
    addStringPart func(int)
    removeStringPart func(int)
    extendStringPart func(int)
}



func (this *Buffer) part() partInterface {
    b := this.sString + this.length 
    if b > ramLength {
       fmt.Println("index out of range")
       // TODO : 2 passes
    }
    a := new(AString)


      this.addStringPart(a.cuid)
    a.cuid = this.cuid
    this.cuid ++ 
    a.rBloc = this.cmp

    a.len = this.length
    a.rStart   = this.sString

    a.rEnd   = b 
    a.rPointer = this.sString
    a.pointer = 0 
    a.strlen = 0
    a.ram = &this.ramString

    a.extend = func(){
        this.sString += this.length  //TODO :: IF NOT USE A OTHER BLOCK
        end := a.rEnd + this.length 
        a.len = a.len   + this.length

        if(this.partString[ a.rBloc + 1 ] == 0){
             if (end > ramLength) {
                    fmt.Println("out of memory")
                    // TODO : 2 passes
             }
             this.extendStringPart(a.cuid)
             a.rBloc ++
             fmt.Println("extend")
             a.rEnd = end 
        }else { 
              
             endForce := this.sString + a.len
             
             if (endForce > ramLength) {
                    fmt.Println("out of memory")
                    // TODO : 2 passes
             }

             copy(this.ramString[this.sString:],this.ramString[a.rStart:a.rEnd]) 
        
             fmt.Println("force extend")

             this.removeStringPart(a.cuid)
             a.rStart   = this.sString
             a.rPointer = this.sString + a.pointer
             this.addStringPart(a.cuid)
             for i:=1 ; i < end / this.length ; i++ {
                    this.extendStringPart(a.cuid)
                  
             }
             a.rBloc    = this.cmp
             a.rEnd     = endForce
             fmt.Println(string(this.ramString[a.rStart:a.rEnd]) )
        }   

    }
  
    this.sString = b 
    return (a)
}


func buffer() BufferInterface{
      
        a:=  new(Buffer)
        a.length        =  ramBloc
        a.ramString     = [ramLength]byte{}

        a.sString       = 0
        a.sInt          = 0
        a.sPartString   = 0
        a.partString    = [nbBloc]int{}
        a.cmp = 0
        a.cuid = 1

        a.extendStringPart = func(cuid int){
            id := a.cmp
            for i,v := range a.partString {
                if v == cuid {  
                    id = i
                }  
            }
            if id == a.cmp {
                a.cmp ++ 
               
            }
            id++
            a.partString[id] = cuid
            fmt.Println(a.partString)
        }

        a.addStringPart = func (cuid int) { 
             a.cmp += space
             a.partString[a.cmp] = cuid
             
             a.sString +=  space*ramBloc
             fmt.Println(a.partString)
        }   

        a.removeStringPart = func (cuid int){
            for i,v := range a.partString {
                if v == cuid {  
                    a.partString[i] = 0
                }  
             }
        }
        
        return (a)
}

///////////////////////////////////////

func main() {

    a:= buffer()

    b:= a.part()
    b.push(`b1`)
    b.push("b2")
    b.push("b3")
    
    c:= a.part()
    c.push("c1")

    b.push("b4")
    c.push("c2")
    b.push(`b5`)

    c.push("c3")
    b.push(`b6`)
    b.push(`b7`)
    c.push("c4")
    b.push(`b7`)
    b.push("bc4")


   
}