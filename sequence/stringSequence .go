package main
////////////////////////
import "fmt"

type Lstring struct {
    table []string
    pointer int
    length int
}

type stringList interface {
    Next() *Lstring 
    Prev() *Lstring
    Print() (int,string)
    Index(int) *Lstring
}

func (this *Lstring) Index(i int) *Lstring{
    if i < this.length && i > 0 {
         this.pointer = i
    } 
    return this
}

func (this *Lstring) Next() *Lstring{
    if this.pointer == this.length -1 {
        fmt.Println("new cycle")
        this.pointer = 0
    } else {
        this.pointer  ++
    }
    return  this 
}

func (this *Lstring) Prev() *Lstring{
    if this.pointer != 0 {
        this.pointer --
    } else {
        fmt.Println("new cycle")
        this.pointer = this.length - 1
    }
    return  this 
}

func (this *Lstring) Print() (int,string){
    return  this.pointer , this.table[this.pointer]
}

func SequenceString(args ...interface{}) stringList {
    a := new(Lstring)
    for _, item := range args {
        switch v := item.(type) {
            case string:
                 a.table = append(a.table, v)
               
        }
    }
    a.length = len(a.table)
    a.pointer = 0  
    return stringList(a) 
}
////////////////////////
func main() {
    a := SequenceString("C5","64","a","c")
    i , c := a.Print()
    fmt.Println(c) 
    fmt.Println(i) 
    _ , c2 := a.Next().Next().Print()
    fmt.Println(c2) 
    _ , c3 := a.Index(3).Print()
    fmt.Println(c3) 
}
