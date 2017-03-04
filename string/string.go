package main
////////////////////
import "fmt"

////////////////////////////////
func Map( callback func(rune)(string),val string) string{

     var start int = 0
     var pt int = 0
     var max int = len(val) 
     /*--------------------------------------------------------------
     *  Other possibility : 
     *  length := 0
     *  max := len(val)*2 ==> hard allocation
     ---------------------------------------------------------------*/
     var w int = 0
     t:= make([]byte, max)
     add := func (char string) {
        /*--------------------------------------------------------------
        *  length + = len(char)
        *  if start + length > max {
        *        t =  append(t,make([]byte, max + len(val)*2)   
        *  }
        ---------------------------------------------------------------*/
        t =  append(t,make([]byte, max + len(char))...)   
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

func Replace( a []rune, b []string ,s string)string {
    /************** 
    /* TODO : no hard allocation (args)
    **************/
    l :=  len(s)
    tmp := []int{}
    for i, c := range s {
        for j, d := range a { 
            if c == d {
                tmp = append(tmp,i)
                tmp = append(tmp,j)
                l += len(d)
            }
        }
    }
    t:= make([]byte, l)
    var w int = 0
    var m int = 0 
    for i := 0; i < len(tmp); i+=2 {
         w += copy(t[w:],s[m:tmp[i]]) 
         w += copy(t[w:],b[tmp[i+1]])
         m = tmp[i]
         m ++
    }
    w += copy(t[w:],s[m:len(s)])
    return  string(t[0:w])
    
}


func EscapeHtml(s string)string {
    a :=  len(s)
    /***************
    * Opti Case ==> hard allocation with a Tampon (dynamic int slice)
    *****************/
    tmp := []int{}
    val := [5]string{"&#34;","&amp;","&#39;","&lt;","&gt;"}
    for i, c := range s {
        switch  c {
                    case 34 : 
                       tmp = append(tmp,i)
                       tmp = append(tmp,0)
                       a += 5
                    case 38 : 
                       tmp = append(tmp,i) 
                       tmp = append(tmp,1)
                       a += 5
                    case 92 : 
                       tmp = append(tmp,i) 
                       tmp = append(tmp,2) 
                       a += 5
                    case 60 :
                       tmp = append(tmp,i) 
                       tmp = append(tmp,3) 
                       a += 4
                    case 62 : 
                       tmp = append(tmp,i) 
                       tmp = append(tmp,4) 
                       a += 4
        }
    }
    t:= make([]byte, a)
    var w int = 0
    var m int = 0 
    for i := 0; i < len(tmp); i+=2 {
         w += copy(t[w:],s[m:tmp[i]])
         w += copy(t[w:],val[tmp[i+1]])
         m = tmp[i]
         m ++
    }
    w += copy(t[w:],s[m:len(s)])
    return  string(t[0:w])
}  
////////////////////////////////

func main() {
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
      a := Map( escapeHtml ," < a > oink & oink & oink") 
      fmt.Println(a) 
      c := Replace( []rune{34,38,92,60,62} ,[]string{"&#34;","&amp;","&#39;","&lt;","&gt;"} ," < a > oink & oink & oink") 
      fmt.Println(c)
      fmt.Println(EscapeHtml(" < a > oink & oink & oink")) 
  
}