package main
import (
    "fmt"
    "crypto/rand"
    "time"
    "encoding/base64"
    "log"
    "net/http"
)


/////////////////////////
func uuid() (string ,error) {
    b := make([]byte, 32)
    _, err := rand.Read(b)
    return  base64.StdEncoding.EncodeToString(b),err
}
//////////////////////////

type ClientInterface interface {
   log()
}

type Client struct{
    roles []int
    expire time.Time
    name string
}

func (this *Client) log() {
    fmt.Println( this.name)
}

/////////////////////

type SessionManagerInterface interface {
    start( http.ResponseWriter, *http.Request) ClientInterface
}

type SessionManager struct {
    clients map[string]*Client
    cookieName string
}

func (this *SessionManager) start(w http.ResponseWriter,r *http.Request) ClientInterface{

   var newCookie func() string
   newCookie = func() string{
        id , _ := uuid()
        expiration := time.Now().Add(365 * 24 * time.Hour)
        cookie := http.Cookie{Name: this.cookieName, Value: id, Expires: expiration}
        http.SetCookie(w, &cookie)
        client := new(Client)
        client.name = "anonymous"
        if  this.clients[id] != nil {
            return newCookie()
        }
        this.clients[id] = client
        return id
    }   
    
    var cookie string = ""
    cookieClient, err := r.Cookie(this.cookieName)

    if err != nil {
        cookie = newCookie()
    }else{
        cookie = cookieClient.Value
        if  this.clients[cookie] == nil {
            cookie = newCookie()
        }
    }    

    return this.clients[cookie]
}

func  sessionManager(cookieName string) SessionManagerInterface{
    a := new(SessionManager)
    a.clients = map[string]*Client{}
    a.cookieName = cookieName
    return a
}

/////////////////////

func main() {
   
    session := sessionManager("GoSess")
    home := func(w http.ResponseWriter, r *http.Request){
        user := session.start(w,r);
        user.log()
        fmt.Fprintf(w, "<html><head></head><body>Welcome</body></html>")
    }
    http.HandleFunc("/",home)
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}

