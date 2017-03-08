package main
////////////////////////
import "fmt"
////////////////////////////////
////////////////////////////////
type stringIterator struct {
    table stringArray
    pointer int
}

func (this *stringIterator) next() (bool,string){
    done := false
    if this.pointer == this.table.length() -1 {
        done = true
        this.pointer = 0
    } else {
        this.pointer  ++
    }
    return  done ,  this.table.get(this.pointer)
}

func (this *stringIterator) Index(i int) *stringIterator{
    if i < this.table.length() && i > 0 {
         this.pointer = i
    } 
    return this
}

func (this *stringIterator) Print() (int,string){
    return  this.pointer , this.table.get(this.pointer)
}

func (this *stringIterator) Next() *stringIterator{
    if this.pointer == this.table.length() -1 {
        fmt.Println("new cycle")
        this.pointer = 0
    } else {
        this.pointer  ++
    }
    return  this 
}

func (this *stringIterator) Done() bool{
    return  this.pointer == 0
}

func (this *stringIterator) Prev() *stringIterator{
    if this.pointer != 0 {
        this.pointer --
    } else {
        fmt.Println("new cycle")
        this.pointer = this.table.length() - 1
    }
    return  this 
}

////////////////////////////////
////////////////////////////////

type AString []string

type stringArray interface {

    pop() string
    shift() string
    push(args ...interface{}) int
    unshift(args ...interface{}) int

    get(int) string
    indexOf(string) int
    join(string) string
    length() int
    values() stringIterator

    Map(func(mapType)string) *AString
    Filter(func(filterType)bool) *AString
    
}
////////////////////////////////
type mapType struct{
    i int
    val string
    table []string
}

type filterType struct{
    i int
    val string
    table []string
}
////////////////////////////////
func (this *AString) indexOf(a string) int{
    for i,v := range *this {
        if v == a {
            return i
        }  
     }
    return -1 
}
func (this *AString) length() int{
    a := *this
    return len(a)
}

func (this *AString) get(index int) string{
    a := *this
    return a[index]
}

func (this *AString) values() stringIterator{
   
    return stringIterator{table:this,pointer:0} 
}

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
////////////////////////////////
func (this *AString) pop() string{
     a := *this
     var d string = a[len(a)-1]
     *this = a[:len(a)-1]
     return d
}

func (this *AString) shift() string{
     a := *this
     var d string = a[0]
     *this = a[1:]
     return d
}

func (this *AString) unshift(args ...interface{}) int{
    r := []string{}
    for _, item := range args {
         switch v := item.(type) {
             case string:
                  r = append(r, v)
         }
     }
     *this = append( r,*this...)
     return len(*this)
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
func (this *AString) Map(a func(mapType)string) *AString{
     d := *this
     for i,v := range d {
	  d[i] = a(mapType{i:i,val:v,table:d})
     }
     return this
}

func (this *AString) Filter(a func(filterType)bool) *AString{
     d := []string{}
     for i,v := range *this {
          if a(filterType{table:d,val:v,i:i}) {
             d = append(d,v)
          }
     }
     *this = d
     return this
}
////////////////////////////////
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

      d1 := a.values()
     
       
      r1 := a.push("c","e")
      a.unshift("45","21")
      fmt.Println(d1.next())
      _ , c2 := d1.Next().Next().Print()
      fmt.Println(c2) 
      _ , c3 := d1.Index(1).Print()
      fmt.Println(c3) 


      fmt.Println(a)
      r3 := a.shift()
      fmt.Println(r3)
      fmt.Println(a)

      fmt.Println(a.indexOf("e"))
      fmt.Println(r1)
      fmt.Println(a.join(","))
      r2 := a.pop()
      fmt.Println(r2)
      fmt.Println(a) 

      c := func(this mapType)string{
      	return this.val + "_super"
      }
      d := func(this filterType)bool{
      	return this.val == "a_super"
      }

      a.Map(c)
      fmt.Println(a)
      a.Filter(d)
    
      fmt.Println(a)
}