package main

import "fmt"
import "time"
import "sort"


////////////////////////
// Array strategy
// Array fixe prediction [1/3]
// No strategy [<3]
////////////////////////

const ramLength = 6
const strategyLength = 3
const ramStrategy = ramLength / strategyLength

///////////////////////////////////////

type ArrayNumericInterface interface {
  push(int)
  Reduce(func(int,int)int,int) int
  Some(func(int)bool) bool
  Sort(func(int,int)int)
}

type AInt struct {
    rEnd int
    pStart int
    pEnd int

    len int
    strlen int

    ram [ramLength]int
}


func (this *AInt) Sort(a func(int,int)int){
    r := map[int]int{}
    b := make([]int,this.len)
    
    j :=  0
    var cmp int
    for i:=this.pStart ; i < this.pEnd ; i++ {
        cmp = 0
        for j:=this.pStart ; j < this.pEnd ; j++ {
            if j != i {
                if a(this.ram[i],this.ram[j]) >= 0 {
                    cmp++
                }
            }
        }
        r[cmp] = this.ram[i]
        b[j] = cmp
        j ++
     }
     c :=  make([]int,j)
     sort.Ints(b)
     for _, k := range b {
        c[k] = r[k]
     }
     copy(this.ram[this.pStart:],c)
     fmt.Println(this.ram) 

}


func (this *AInt) Some(a func(int)bool) bool{
     for i:=this.pStart ; i < this.pEnd ; i++ {  
        if a(this.ram[i]) {
            return true
        }
    }
    return false
}


func (this *AInt) Reduce(a func(int,int)int,b int) int{
    var r int = b
    for i:=this.pStart ; i < this.pEnd ; i++ {   
	 r +=  a(r, this.ram[i])
    }
    return r
}

func (this *AInt) push(v int)  {
    if this.pEnd  >= this.rEnd {  
            s:=  this.pStart
            if s != 0 {   // Move the head
                fmt.Println( "Move the head push") 
                d := s
                if s > strategyLength {
                    d = s/strategyLength 
                } else {
                    d = 0
                }
                copy(this.ram[d:],this.ram[this.pStart:this.pEnd])
                this.pEnd -= s - d
                this.pStart = d

            }else {
                fmt.Println( "Out of memory") 
            } 
    }
   
    this.ram[this.pEnd] = v
    this.pEnd++    
    this.len++
}

///////////////////////////////////////

type ArrayStringInterface interface {
    unshift(string)
    push(string) 
    
    shift() string
    pop() string

    indexOf(string) int
    join(string) string

    length() int
    innerLength() int

    Map(func(mapStringType)string) 
    Filter(func(mapStringType)bool) 

}

type AString struct {
    rEnd int
    pStart int
    pEnd int
   
    len int
    strlen int

    ram [ramLength]string
}

func (this *AString) length() int  {
    return this.len
}

func (this *AString) innerLength() int  {
    return this.strlen
}

func (this *AString) shift() string  {
     if this.pStart < this.pEnd {
        this.pStart++
        this.len--
        v := this.ram[this.pStart-1]
        this.strlen -= len(v)
        return v
    }else {
        return "@nil"
    }
}

func (this *AString) pop() string  {
     if this.pEnd > this.pStart {
         this.pEnd--
         this.len--
         v := this.ram[this.pEnd]
         this.strlen -= len(v)
         return v
     }else {
        return "@nil"
     }
}

func (this *AString) push(v string)  {
    if this.pEnd  >= this.rEnd {  
            s:=  this.pStart
            if s != 0 {   // Move the head
                fmt.Println( "Move the head push") 
                d := s
                if s > strategyLength {
                    d = s/strategyLength 
                } else {
                    d = 0
                }
                copy(this.ram[d:],this.ram[this.pStart:this.pEnd])
                this.pEnd -= s - d
                this.pStart = d

            }else {
                fmt.Println( "Out of memory") 
            } 
    }
    this.strlen += len(v)
    this.ram[this.pEnd] = v
    this.pEnd++    
    this.len++
}

func (this *AString) unshift(v string)  {
    if(this.pStart <= 0){ 
            s:=  this.rEnd -  this.pEnd 
            if s != 0 {   // Move the head
                fmt.Println( "Move the head unshift") 
                if s > strategyLength {
                    s = s/strategyLength
                }
                copy(this.ram[ s:],this.ram[0:this.pEnd])
         
                this.pStart =  s
                fmt.Println(this.pStart) 
                this.pEnd += s
            }else {
                fmt.Println( "Out of memory") 
            } 
    }
    this.pStart--
    this.len++
    this.strlen += len(v)
    this.ram[this.pStart] = v
}


func (this *AString) join(a string) string {   
     /// 247127700
     t:= make([]byte, this.strlen + (this.len * len(a)))
     j:=0
     if this.len != 0 {
         j +=  copy(t[j:],this.ram[this.pStart])
         for i:=this.pStart+1 ; i < this.len ; i++ {
                  j +=  copy(t[j:],a)
                  j +=  copy(t[j:],this.ram[i])

         }
    }
    return  string(t[:j])
}


func (this *AString) indexOf(a string) int{
     for i:=this.pStart ; i < this.pEnd ; i++ {
        if this.ram[i] == a {
            return i
        }  
     }
    return -1 
}

////////////////////////////////
type mapStringType struct{
    i int
    val string
}

func (this *AString) Map(a func(mapStringType)string){
     for i:=this.pStart ; i < this.pEnd ; i++ {   
          t1 := this.ram[i]
          this.strlen -= len(t1)
          t2 := a(mapStringType{i:i,val:this.ram[i]})
          this.strlen += len(t2)
	  this.ram[i] = t2
     }
}

func (this *AString) Filter(a func(mapStringType)bool){
     t:= make([]string, this.len)
     j:=0
     for i:=this.pStart ; i < this.pEnd ; i++ {   
          if a(mapStringType{i:i,val:this.ram[i]}) {
                t[j] = this.ram[i]
                j ++ 
          }
     }
    
     copy(this.ram[this.pStart:],t[0:j])
     this.pEnd = this.pStart + j 
     this.len = j
}



///////////////////////////////////////


func ArrayStringBuilder() ArrayStringInterface{
       
    a := new(AString)

    a.len      = 0
    a.rEnd     = ramLength
    a.pEnd = ramStrategy
    a.pStart   = ramStrategy
    a.strlen   = 0
    a.ram      = [ramLength]string{}

    return (a)
}


func ArrayIntBuilder() ArrayNumericInterface{
       
    a := new(AInt)

    a.len      = 0
    a.rEnd     = ramLength
    a.pEnd     = ramStrategy
    a.pStart   = ramStrategy
    a.ram      = [ramLength]int{}

    return (a)
}



///////////////////////////////////////
//  20514000
///////////////////////////////////////
func main() {

    t1 := time.Now().UnixNano()


    fmt.Println("-------NUMERIC------")
    a:= ArrayIntBuilder()
    a.push(4)
    a.push(3)

    fmt.Println("-------REDUCE------")
    c1 := a.Reduce(func(c int,v int)int{
      	 return c + v
    },0)
    fmt.Println(c1) 
    
    c2 := a.Some(func(c int) bool {
      	 return c > 3
    })
    fmt.Println(c2) 
    fmt.Println("-------SORT------")
    a.Sort(func(c int,v int)int{
      	 return c - v
    })
    
    b:= ArrayStringBuilder()

    b.push("p1")
    b.push("p2")
    b.push("p3")
    b.push("p4")
    b.push("p5")

   // b.push("hello world")

    b.unshift("u1")
    
    c := b.shift()
    b.unshift("u2")


    fmt.Println(b.join(" , ")) 
    fmt.Println(b.indexOf("p5")) 

    mapS := func(this mapStringType)string{
      	return this.val + "_#"
    }

    b.Map(mapS)
    fmt.Println(b.join(" , ")) 
    fmt.Println("-------INDEX------")
    fmt.Println(b.indexOf("u2_#")) 
    
    fmt.Println("-------FILTER------")
    filterS := func(this mapStringType)bool{
      	return this.val == "u2_#" || this.val == "p2_#" || this.val == "p3_#"
    }
    b.Filter(filterS)

    fmt.Println(b.join(" , ")) 
    fmt.Println("-------LENGTH------")
    fmt.Println(b.length()) 
    fmt.Println(b.innerLength()) 
    fmt.Println("-------POP------")
    c = b.pop()
    fmt.Println(c)
    c = b.pop()
    fmt.Println(c)
    fmt.Println("------SHIFT-----")

    c = b.shift()
    fmt.Println(c)
    c = b.shift()
    fmt.Println(c)
    c = b.shift()
    fmt.Println(c)
    c = b.shift()
    fmt.Println(c)
    c = b.shift()
    fmt.Println(c)
    c = b.shift()
    fmt.Println(c)
    c = b.shift()
    fmt.Println(c)

    fmt.Println("-------------")
    fmt.Println(b.join(" , ")) 
    fmt.Println("-------------")
   
    fmt.Println(time.Now().UnixNano()-t1) 


}