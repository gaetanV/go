package main
////////////////////////
import "fmt"
import "time"

const ramLength = 50000
var ram [ramLength]byte = [ramLength]byte{}
var s1 int = 0 

type ArrayStringInterface interface {
    push(string) int
}

type AString struct {
    rStart int
    rEnd int
    rPointer int
    step int
    len int
    pointer int
}

func (this *AString) push(v string) int {
    if this.rPointer + len(v) >= this.rEnd {      
        length :=  this.len + this.step
        this.len = length
        if s1 != this.rEnd {
            b := s1 + length
            this.rPointer = s1 + this.pointer
            if b > ramLength {
                s1 = 0
            }
            end := b    
            copy(ram[s1:],ram[this.rStart:this.rEnd]) 
            this.rStart = s1 
            this.rEnd = end
            s1 = end
        }else {
            this.rEnd = this.rEnd + length
            s1 = this.rEnd
        }
    }
    this.rPointer += copy(ram[this.rPointer:],v) 
    this.pointer ++     
    return this.len 
}

func ArrayString(length  int) ArrayStringInterface {
    b := s1 + length 
    if b > ramLength {
        s1 = 0
    }
    a := new(AString)
    a.rStart   = s1
    a.rEnd   = b
    a.step = length
    a.len = length
    a.rPointer = s1
    a.pointer = 0 
    s1 = b
    return (a)
}


////////////////////////
// 3079162600 Args
// 557451100  String
///////////////////////

func main() {

    var t1 int64
    limit := 2000000
    var c1 int64 = 0

    var i int

    i = 0
    for { 

        if i >= limit { break }  
        i ++
        t1 = time.Now().UnixNano()
        a := ArrayString(25)
        a.push("caa")
        a.push("cadda")
        a.push("daa")

        c := ArrayString(3)
        c.push("aa")
        c.push("add")
        c.push("a")

        b := ArrayString(2)
        b.push("c")
        b.push("c")
        b.push("d")
        b.push("e")
        b.push("e")
        b.push("e")    
        c.push("e")
        c.push("aaa")
        c.push("e")    
        c.push("e")
        c.push("e")
        c.push("e")
        c1 += time.Now().UnixNano()-t1
    } 
    fmt.Println(c1) 
    fmt.Println(ram[0:1000])
}