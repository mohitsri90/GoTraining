package main

import "fmt"

type purchase interface {
	sell()
}

type display interface {
	show()
}

//we use two previous interfaces to form the following composite interface through embedding
type salesman interface {
	purchase
	display
}

type game struct {
	name, price    string
	gameCollection []string
}

//we use game struct to implement interfaces
func (t game) sell() {
	fmt.Println("----------------------------")
	fmt.Println("Name: ", t.name)
	fmt.Println("Price: ", t.price)
	fmt.Println("----------------------------")
}

func (t game) show() {
	fmt.Println("The Games available are: ")
	for _, name := range t.gameCollection {
		fmt.Println(name)
	}
	fmt.Println("---------------------------")
}

//this method takes composite interface as parameter
//Since the interface is composed of purchase and display
//Hence the child methods of those interfaces can be accessed here
func shop(s salesman) {
	fmt.Println(s)
	s.sell()
	s.show()
}

func main() {
	collection := []string{"XYZ", "Trial by code", "Sea of rubble"}
	game1 := game{"ABC", "$125", collection}
	shop(game1)
}
