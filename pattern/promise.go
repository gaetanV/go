package main
import "fmt"
////////////////////////
type Promise struct{
     then func(func(string),func(string)) 
}
func newPromise(a func( promiseType)) Promise{
    return Promise{then:func(resolve func(string) , reject func(string)){ 
        success := func(a string){
            resolve(a)   
        }
        error := func(a string){
            reject(a)
        }
        a(promiseType{resolve:success,reject:error})
    }}
}
////////////////////////////////
type promiseType struct{
    resolve func(string)
    reject func(string) 
}
////////////////////////
func main() {
     a := func(this promiseType) {
          this.resolve("succes")
         // this.reject("error")
     }
     b := func(resolve string){
          fmt.Println(resolve)
     }
     c := func(reject string){
          fmt.Println(reject)
     }
     newPromise(a).then(b,c)
}