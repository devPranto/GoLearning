package main

import (
	"fmt"
	//"strings"
)
type car interface {
	spec()
}
type sportsCar struct {
	model string
	HP    int
	brand string
}
type sedan struct{
	model string 
	seats int 
	cargoSpace int 
}
func main() {
	Ferrari := sportsCar{
		model: "Ferrari F12",
		HP:    712,
		brand: "ferrari",
	}

	AMG := sportsCar{
		model: "AMG GT4",
		HP: 800,
		brand: "Merecedes",
	}
	Sonata := sedan{
		model: "Hyundai Sonata",
		seats: 4,
		cargoSpace: 3,
	}

	AMG.spec()
	Ferrari.spec()
	Sonata.spec()
	fmt.Println("==============================")
	print(Sonata)
	print(AMG)

}

func(s sportsCar) spec(){	
	fmt.Println(s)
}
func(s sedan) spec(){
	fmt.Println(s)
}
func print(s car){
	fmt.Println(s)
}