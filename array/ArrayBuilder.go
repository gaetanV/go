package main
////////////////////////
import "fmt"
//import "time"

const ramLength = 50000
var ramString [ramLength]string = [ramLength]string{}
var ramInt [ramLength]int = [ramLength]int{}
var sString int = 0 
var sInt int = 0 


type ArrayIntInterface interface {
  push(int)
}


type AInt struct {
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
            this.rEnd = this.rEnd + length
            *this.follow = this.rEnd
         }
    }
    this.ram[this.rPointer] = v
    this.rPointer++
    this.pointer ++     
}


type ArrayStringInterface interface {
    push(string) 
}

type AString struct {
    rStart int
    rPointer int
    rEnd int
    step int
    len int
    pointer int
    follow *int
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
            this.rEnd = this.rEnd + length
            *this.follow = this.rEnd
         }
    }
    this.ram[this.rPointer] = v
    this.rPointer++
    this.pointer ++     
}

type ArrayBuilder struct {}

func (this *ArrayBuilder) int(length int) ArrayIntInterface {
    b := sInt + length 
    if b > ramLength {
        sInt = 0
    }
    a := new(AInt)
    a.step = length
    a.len = length
    a.rStart   = sInt
    a.rPointer = sInt
    a.pointer = 0 
    a.rEnd   = b
    a.follow = &sInt
    a.ram = &ramInt
    sInt = b
    return (a)
}


func (this *ArrayBuilder) string(length int) ArrayStringInterface {
    b := sString + length 
    if b > ramLength {
        sString = 0
    }
    a := new(AString)
    a.step = length
    a.len = length
    a.rStart   = sString
    a.rPointer = sString
    a.pointer = 0 
    a.rEnd   = b
    a.follow = &sString
    a.ram = &ramString
    sString = b
    return (a)
}

func main() {
    a:= new(ArrayBuilder) 
    d:= a.string(1)
    c:= a.string(1)
    c.push("go c")
    b:= a.string(2)
    b.push("hello world")
    d.push("c++")
    b.push("hello world")
    b.push("hello world")
    c.push("go c")

    d.push("c++")
    d.push("c++")
    

    e:= a.int(1)
    e.push(45)
    e.push(555)
    e.push(2)
    fmt.Println(ramInt[0:50]) 
    fmt.Println(ramString[0:50]) 

}