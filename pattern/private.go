package main
import "fmt"
////////////////////////
type Public interface{
    getPublic() string
}

type Private struct{
     getPrivate func(string) string
     setPrivate func(string,string) string
     public string
}

func (this *Private)privateFunc() string{
    return "private"
}

func (this *Private)getPublic() string{
    fmt.Println(this.getPrivate("wait"))
    this.setPrivate("wait","45")
    fmt.Println(this.getPrivate("wait"))
    return this.public
}

func PrivateBuild() Public{
    var a *Private = new(Private)
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
    return a
}

////////////////////////
func main() {
    var a Public = PrivateBuild()
    fmt.Println(a)
    fmt.Println(a.getPublic())
}
