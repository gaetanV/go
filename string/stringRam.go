package main

import "fmt"
import "time"
////////////////////
///////////////////////////////
const replaceByteExtra = 50
const replaceInt = 8 // => 4 replace

func EscapeHtmlPredict(s string)string {

    // Time               553639800 |  443669500

    max := replaceByteExtra
    t:= make([]byte, len(s) + replaceByteExtra)
    var w int = 0
    var m int = 0 
    add := func(val string,pointer int){
         if pointer + len(val) > max {
              max += replaceByteExtra
              t =  append(t,make([]byte, max)...)   
         } 
         w += copy(t[w:],s[m:pointer])
         w += copy(t[w:],val)
         m = pointer
         m ++
    }
    for i, c := range s {
        switch  c {
                    case 34 : 
                      add("&#34",i)
                    case 38 : 
                      add("&amp;",i)
                    case 92 : 
                      add("&#39;",i)
                    case 60 :
                      add("&lt;",i)
                    case 62 : 
                      add("&gt;",i)
        }
    }
    w += copy(t[w:],s[m:len(s)])
    return  string(t[0:w])
}
///////////////////

func EscapeHtmlBuffer(s string)string {

    // Time         553639800 | 465136400

    a :=  len(s)
    b := replaceInt
    tmp  := make([]int,replaceInt)
    tmp2 := make([]int,replaceInt)
    o := 0
    f := func(i int ,v int, l int ){
             tmp[o] = i
             tmp2[o] = v
             o++
             a += 5
             if o > b {
                 b += replaceInt
                 tmp =   append(tmp,make([]int,b)...)   
                 tmp2 =  append(tmp2,make([]int,b)...)  
             }
    }
    for i := 0 ; i < len(s) ; i++{
        switch  s[i] {
                    case 34 :  
                       f(i,0,5)
                    case 38 : 
                       f(i,1,5)
                    case 92 :
                        f(i,2,5) 
                    case 60 :
                        f(i,3,4) 
                    case 62 : 
                        f(i,4,4) 
        }
    }
    t:= make([]byte, a)
    var w int = 0
    var m int = 0 
    for i := 0; i < o; i++ {
         w += copy(t[w:],s[m:tmp[i]])
         w += copy(t[w:],val[tmp2[i]])
         m = tmp[i]
         m ++
    }
    w += copy(t[w:],s[m:len(s)])
    return  string(t[0:w])
}  
////////////////////////////////


var val [5]string = [5]string{"&#34;","&amp;","&#39;","&lt;","&gt;"}
func MapBuffer( callback func(rune)(string),s string) string{

    //  Time                       834175800 |   727789000            

    a :=  len(s)
    b := replaceInt
    tmp  := make([]int,replaceInt)
    tmp2 := make([]string,replaceInt)  
 
     o := 0
     for i, c := range s {
        d := callback(c)
         if d != "@nil"  {
             tmp[o] = i
             tmp2[o] = d
             o++
             a += len(d)
             if o > b {
                 b += replaceInt
                 tmp =   append(tmp,make([]int,b)...)   
                 tmp2 =  append(tmp2,make([]string,b)...)
             }
        }
     }

    t:= make([]byte, a)
    var w int = 0
    var m int = 0 
    for i := 0; i < o; i++ {
         w += copy(t[w:],s[m:tmp[i]])
         w += copy(t[w:],tmp2[i])
         m = tmp[i]
         m ++
    }
    w += copy(t[w:],s[m:len(s)])
    return  string(t[0:w])
}

///////////////////

func MapPredict( callback func(rune)(string),s string) string{

     //  Time                       834175800 | 571698500       

     max := replaceByteExtra
     t:= make([]byte, len(s) + replaceByteExtra)
     var w int = 0
     var m int = 0 
    
     add := func (val string,pointer int) {
        if pointer + len(val) > max {
              max += replaceByteExtra
              t =  append(t,make([]byte, max)...)   
        } 
        w += copy(t[w:],s[m:pointer])
        w += copy(t[w:],val)
        m = pointer
        m ++
     }
    for i, a := range s {
        d := callback(a)
        if d != "@nil"  {
            add (d,i)
        }
    }
    w += copy(t[w:],s[m:len(s)])
    return  string(t[0:w])
}
///////////////////

func main() {
    var t1 int64
    limit := 2000000
    var c1 int64 = 0
    var c2 int64 = 0
    var c3 int64 = 0
    var c4 int64 = 0
    var i int

    i = 0
    for { 
            if i >= limit { break }  
            i ++
            t1 = time.Now().UnixNano()
            EscapeHtmlBuffer(" < a > oink & oink & oink")
            c1 += time.Now().UnixNano()-t1
             
    } 
    fmt.Println(c1) 

    i = 0
    for { 
            if i >= limit { break }  
            i ++ 
            t1 = time.Now().UnixNano()
            EscapeHtmlPredict(" < a > oink & oink & oink")
            c2 += time.Now().UnixNano()-t1
             
    } 
    fmt.Println(c2) 

    i = 0
    for { 
            if i >= limit { break }  
            i ++
            t1 = time.Now().UnixNano()
            escapeHtml := func(a rune)(string){
                  switch a {
                          case 34 : 
                              return "&#34;"
                          case 38 : 
                              return "&amp;"
                          case 92 : 
                              return "&#39;"
                          case 60 : 
                              return "&lt;"
                          case 62 : 
                              return "&gt;"
                  }
                  return "@nil"
            }
            MapPredict( escapeHtml ," < a > oink & oink & oink")
            c3 += time.Now().UnixNano()-t1
             
    } 
    fmt.Println(c3) 


    i = 0
      for { 
            if i >= limit { break }  
            i ++
            t1 = time.Now().UnixNano()
            escapeHtml := func(a rune)(string){
                  switch a {
                          case 34 : 
                              return "&#34;"
                          case 38 : 
                              return "&amp;"
                          case 92 : 
                              return "&#39;"
                          case 60 : 
                              return "&lt;"
                          case 62 : 
                              return "&gt;"
                  }
                  return "@nil"
            }
            MapBuffer( escapeHtml ," < a > oink & oink & oink")

            c4 += time.Now().UnixNano()-t1
             
    } 
    fmt.Println(c4) 
}