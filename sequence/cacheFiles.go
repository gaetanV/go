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

func threadBoot(file string, resolve  func(int,int64,[]byte) , error func(error) ){
    fileContent, err := os.Open(file)
    if err != nil {
        error(err)
    }
    defer fileContent.Close()
    fi, err := fileContent.Stat()
    if err != nil {
       error(err)
    }
    length := fi.Size()
    buf := make([]byte,  length)
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
    resolve(r,length,buf[0:n])
}

func boot(namespace string, promise func(bootResponse)){
    var nbView int = 0
    var complet int = 0
    var path string = fmt.Sprintf("%s/%s/",repertory,namespace)

    d, err := os.Open(path)
    if err != nil {
        return
    }
    defer d.Close()
    files, _ := d.Readdirnames(-1)
    

    var nbFiles int = len(files)

    tamponIndex   := make([]int,nbFiles)
    tamponContent := make([][]byte,nbFiles)

    next := func(){
        complet++
        if complet >= nbFiles {
            promise(bootResponse{tamponIndex:tamponIndex,tamponContent:tamponContent,nbView:nbView+1})
        }
        
    }
    error := func(err error){
        fmt.Println(err)
        next()
    }
    load := func(index int , length int64, values []byte ){
        if index > nbView {
            nbView = index 
        }
        tamponIndex[complet]   = index
        tamponContent[complet] = values
        next() 
    }

    for i:= 0 ; i < nbFiles ; i++ {
        go threadBoot(path+files[i],load,error)
    }
}

type bootResponse struct {
    tamponIndex   []int
    tamponContent [][]byte
    nbView        int
}

func main() {

    init := func(r bootResponse){
            tmp := make([][]byte,r.nbView + predict)
            for i,data := range r.tamponContent{
                tmp[r.tamponIndex[i]] = data
            }
            fmt.Println(tmp)
    }
    boot("page",init)
    
    time.Sleep(time.Second * 2)

}

