package main
////////////////////////
import "fmt"

type Aint []int

type intArray interface {
   Reduce(func(int,int)int,int) int
}
func (this *Aint) Reduce(a func(int,int)int,b int) int{
    var r int = b
    for _,v := range *this {
	 r +=  a(r,v)
    }
    return r
}


func ArrayInt(args ...interface{}) intArray {
    a := new(Aint)
    for _, item := range args {
        switch v := item.(type) {
            case int:
                 *a = append(*a, v)
        }
    }
    return intArray(a) 
}
////////////////////////
func main() {
      a := ArrayInt(5,6,5456)

      var c int = a.Reduce(func(c int,v int)int{
      	 return c + v
      },0)

      fmt.Println(c) 
}