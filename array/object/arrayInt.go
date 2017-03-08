package main
////////////////////////
import "fmt"
import "sort"

type Aint []int

type intArray interface {
   Reduce(func(int,int)int,int) int
   Some(func(int)bool) bool
   Sort(func(int,int)int)
}

func (this *Aint) Reduce(a func(int,int)int,b int) int{
    var r int = b
    for _,v := range *this {
	 r +=  a(r,v)
    }
    return r
}

func (this *Aint) Some(a func(int)bool) bool{
    for _,v := range *this {
        if a(v) {
            return true
        }
    }
    return false
}

func (this *Aint) Sort(a func(int,int)int){
    r := map[int]int{}
    var cmp int
     for i,v1 := range *this {
        cmp = 0
        for j,v2 := range *this {
            if j != i {
                if a(v1,v2) >= 0 {
                    cmp++
                }
            }
        }
        r[cmp] = v1
     }
    var keys []int
    for k := range r {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    rest := []int{}
    for _, k := range keys {
        rest = append(rest,r[k])
    }
     *this = rest
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
      a := ArrayInt(9879,454,4565,5,4,5456)

      a.Sort(func(c int,v int)int{
      	 return c - v
      })
      fmt.Println(a) 

      c1 := a.Reduce(func(c int,v int)int{
      	 return c + v
      },0)
      fmt.Println(c1) 

      c2 := a.Some(func(c int) bool {
      	 return c > 200
      })
      fmt.Println(c2) 
}