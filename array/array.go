package main
////////////////////////
type Array struct{
	a interface{}
}
func (this *Array) push(a interface{}) *Array{
	switch v := this.a.(type) {
		case []string:
                    switch z := a.(type) {
                        case string:
                            this.a = append(v,z)
                    }
		case []int:
                    switch z := a.(type) {
                        case int:
                            this.a = append(v,z)
                    }	
	}
	return this
} 
func ArrayList(a interface{}) Array {
    return Array {a:a}
}
////////////////////////
func main() {
      a := ArrayList([]string{"v"})
      b := ArrayList([]int{4})
      a.push("5")
      b.push(5)
}