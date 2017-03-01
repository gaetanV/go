package main
////////////////////////
import "fmt"

type AString []string

type stringArray interface {
    join(string) string
    pop() string
    push(args ...interface{}) int
    Map(func(int,string)string) *AString
    Filter(func(int,string)bool) *AString
}

////////////////////////////////
func (this *AString) pop() string{
     a := *this
     var d string = a[len(a)-1]
     *this = a[:len(a)-1]
     return d
}

func (this *AString) push(args ...interface{}) int{
    for _, item := range args {
         switch v := item.(type) {
             case string:
                  *this = append(*this, v)
         }
     }
     return len(*this)
}
////////////////////////////////
func (this *AString) Map(a func(int,string)string) *AString{
     d := *this
     for i,v := range d {
	  d[i] = a(i,v)
     }
     return this
}

func (this *AString) Filter(a func(int,string)bool) *AString{
     d := []string{}
     fmt.Println(d)
     for i,v := range *this {
          if a(i,v) {
             d = append(d,v)
          }
     }
     *this = d
     return this
}
////////////////////////////////

func (this *AString) join(a string) string{
     var r string = ""
     for i,v := range *this {
          if i != 0 {
              r += a 
          }
          r += v
     }
    return r
}

func ArrayString(args ...interface{}) stringArray {
    a := new(AString)
    for _, item := range args {
        switch v := item.(type) {
            case string:
                 *a = append(*a, v)
        }
    }
    return stringArray(a) 
}
////////////////////////
func main() {
      a := ArrayString("a","b")
      r1 := a.push("c","e")

      fmt.Println(r1)
      fmt.Println(a.join(","))
      r2 := a.pop()
      fmt.Println(r2)
      fmt.Println(a) 

      c := func(i int,v string)string{
      	return v + "_super"
      }
      d := func(i int,v string)bool{
      	return v == "a_super"
      }
      fmt.Println(a)
      a.Map(c).Filter(d)

    
}