package main
////////////////////////
import "fmt"

type AString []string

type stringArray interface {
    push(string) *AString
}

func (this *AString) push(a string) *AString{
     *this = append(*this,a)
     return this
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
      a.push("c");
      fmt.Println(a)
}