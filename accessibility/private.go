package main
import "fmt"
////////////////////////
type Private interface{
    getPublic() string
}

type PrivateClass struct{
     getPrivate func(string) string
     setPrivate func(string,string) string
     public string
}

func (this *PrivateClass)privateFunc() string{
    return "private"
}

func (this *PrivateClass)getPublic() string{
    fmt.Println(this.getPrivate("wait"))
    this.setPrivate("wait","45")
    fmt.Println(this.getPrivate("wait"))
    return this.public
}

func PrivateBuild() Private{
    a := new(PrivateClass)
    b := map[string]string{}
    b ["wait"] = "privateKey"

    getPrivate := func(i string) string {
       return b[i]
    }
    setPrivate := func(i string, value string) string {
       b[i] = value
       return b[i]
    }

    a.getPrivate = getPrivate
    a.setPrivate = setPrivate

    a.public = "publicKey"
    return  a
}

////////////////////////
func main() {
    a := PrivateBuild()
    fmt.Println(a)
    fmt.Println(a.getPublic())
}
