package main

import "fmt"

type person struct {
	name     string
	contact contactInfo
	id       int
	password int
}
type contactInfo struct{
	email string
	zipCode int
}
func (p person) print(){
	fmt.Printf("%+v\n",p)
}

func (p *person) updateName(name string){
	(*p).name=name
}

func main() {
	dev := person{
		name:     "Pranto",
		contact:contactInfo{email:  "prantodev1@gmail.com",
							zipCode: 1100}  ,
		id:       1,
		password: 1424}
	devPointer:=&dev
	devPointer.updateName("Pranto Dev")
	dev.print()
}
