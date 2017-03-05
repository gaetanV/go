package main
////////////////////
import "fmt"
import "time"
////////////////////////////////
func Map( callback func(rune)(string),val string) string{

    //  Time                       834175800       

     var start int = 0
     var pt int = 0
     var max int = len(val)
     var w int = 0
       for _, a := range val {
       d := callback(a)
        if d != "@nil"  {
          max += len(d)
       }
      }
     t:= make([]byte, max)
     add := func (char string) {
        w += copy(t[w:],val[pt:start])
        w += copy(t[w:],char)
        pt = start + 1
     }
    for _, a := range val {
        d := callback(a)
        if d != "@nil"  {
            add (d)
        }
        start ++
    }
    w += copy(t[w:],val[pt:start])
    return  string(t[0:w])
}
func EscapeHtml(s string)string {

    // Time               553639800 

    a :=  len(s)
    for _, c := range s {
        switch  c {
                    case 34 : 
                       a += 5
                    case 38 : 
                       a += 5
                    case 92 : 
                       a += 5
                    case 60 :
                       a += 4
                    case 62 : 
                       a += 4
        }
    }
    t:= make([]byte, a)
    var w int = 0
    var m int = 0 
    add := func(val string,pointer int){
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
////////////////////////////////

func main() {
      var t1 int64
      limit := 2000000
      var c1 int64 = 0
      var c2 int64 = 0
      var i int

    i = 0
    for { 
            if i >= limit { break }  
            i ++
            t1 = time.Now().UnixNano()
            EscapeHtml(" < a > oink & oink & oink")
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
            Map( escapeHtml ," < a > oink & oink & oink")
            c1 += time.Now().UnixNano()-t1
             
    } 
    fmt.Println(c1) 
}