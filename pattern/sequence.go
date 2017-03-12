package main
////////////////////////
import "fmt"

type PublicList interface{
     Run() bool
}

type PrivateList struct{
     funcRun func() 
     nextPointer func() int
}

func (this *PrivateList) Run()bool {
    this.funcRun()
    i := this.nextPointer()
    if i == 0 {
        return true
    }else { 
        return false
    }
}

func Sequence(args ...interface{}) PublicList {
    a := new(PrivateList)
    table := []func(){}
    for _, item := range args {
        switch v := item.(type) {
            case func():
                table = append(table, v)
               
        }
    }
    pointer := 0
    length := len(table)
    a.funcRun = func()  {
        table[pointer]()
    }

    a.nextPointer = func() int{
        if pointer >= length -1 {
           pointer = 0 
        }else{
          pointer ++
        }
        return pointer
    }   
    return a 
}

////////////////////////
func main() {
    var new string = "new cycle"

    f1 := func (){
        fmt.Println("func1")
    }
    f2 := func (){
        fmt.Println("func2")
    }
    a := Sequence(f1,f2)
     
    if a.Run() {
        fmt.Println(new)
    }
    if a.Run() {
        fmt.Println(new)
    }
    if a.Run() {
        fmt.Println(new)
    }
}
