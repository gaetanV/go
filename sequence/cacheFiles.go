package main
import (
    "fmt"
    "time"
    "os"
    "unicode"

)

/////////////////////////
const repertory = "cache"
const predict = 5

func threadBoot(file string, resolve  func(int,[]byte), error func(error) ){
    fileContent, err := os.Open(file)
    if err != nil {
        error(err)
    }
    defer fileContent.Close()
    fi, err := fileContent.Stat()
    if err != nil {
       error(err)
    }
    buf := make([]byte, fi.Size())
    n, err :=  fileContent.Read(buf)
    if err != nil {
        error(err)
    }
    name := fi.Name()
    var r  int = 0  
    for i:=0 ; i < len(name); i++ {
        if unicode.IsDigit(rune(name[i])) {                                                                                  
            r = r*10 + int(name[i]-'0')                                                                                   
        } else {
            break
        }    
    }
    resolve(r,buf[0:n])
}

func boot(namespace string, promise func(CacheInterface)) {

    var nbView int = 0
    var complet int = 0
    var path string = fmt.Sprintf("%s/%s/", repertory, namespace)

    d, err := os.Open(path)
    if err != nil {
        return
    }
    defer d.Close()
    files, _ := d.Readdirnames(-1)

    var nbFiles int = len(files)

    tamponIndex   := make([]int, nbFiles)
    tamponContent := make([][]byte, nbFiles)

    next := func() {
        complet++
        if complet >= nbFiles {
            tmp := make([][]byte,nbView + predict)
            for i,data := range tamponContent {
                tmp[tamponIndex[i]] = data
            }
            promise(cache(namespace,tmp))
        }
        
    }
    error := func(err error) {
        fmt.Println(err)
        next()
    }
    load := func(index int, values []byte ) {
        if index > nbView {
            nbView = index 
        }
        tamponIndex[complet] = index
        tamponContent[complet] = values
        next() 
    }

    for i:= 0 ; i < nbFiles ; i++ {
        go threadBoot(path+files[i], load, error)
    }
}




////////////////////////

type CacheInterface interface{
    get(int) string
    set(int,*Buffer,func(),func(error))  int
}

type Cache struct {
    memoryContent [][]byte
    namespace string
}

func (this *Cache) get(id int) string {
    return string(this.memoryContent[id])
}

func (this *Cache) set(id int, buffer *Buffer , resolve func(), error func(error)) int {
    this.memoryContent[id] = buffer.tampon
    path := fmt.Sprintf("%s/%s/%d", repertory, this.namespace, id)
    f, err := os.Create(path)
    if err != nil {
        error(err)
    }
    fmt.Println(f.Fd())
    defer f.Close()
    cmp, err :=  f.Write(buffer.tampon)
    if err != nil {
       error(err)
    }
    f.Sync()
    resolve()
    return cmp
}


func cache(namespace string, memory [][]byte) CacheInterface{
    a:= new(Cache)
    a.memoryContent = memory
    a.namespace = namespace 
    return a
}

type Buffer struct {
    tampon []byte
}

func get(pages CacheInterface) {
   pages.get(3)
   fmt.Println(pages.get(3))   
    
}
func set(pages CacheInterface, i int) {
    buffer := new(Buffer)
    t :=  make([]byte,3)
    t[0] = byte(i+32)
    buffer.tampon = t

    success := func(){
         fmt.Println(pages.get(3))
    }
    error := func(err error) {
         fmt.Println(err)
    }

    pages.set(3,buffer,success , error) 
}

func main() {
    init := func(pages CacheInterface) {
        fmt.Println(pages.get(2))   
           
        for i:=0;  i<1000 ; i++ {
            go get(pages)
        }
        for i:=0;  i<100 ; i++ {
            go set(pages,i)
        }    
    }
    boot("page",init)
    time.Sleep(time.Second * 10)
}

