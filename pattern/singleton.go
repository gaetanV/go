package main
////////////////////////
type Peoples struct{
    list map[int]*People 
}
func (this *Peoples) getPeople(id int) *People{
    value, ok := this.list[id]
    if !ok {
       value = newPeople(id)
       this.list[id] = value
    }
    return value
}
////////////////////////
type People struct {
    id  int
}
func newPeople(id int) *People{
    return &People{id:id}
}
////////////////////////
func main() {
    peoples := Peoples{list:map[int]People{}}
    people1 := peoples.getPeople(1)
    people2 := peoples.getPeople(2)
}