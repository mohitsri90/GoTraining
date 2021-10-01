package main

import "fmt"

func add_num(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Result: ", res)
	return res
}

//main fn

func main() {

	fmt.Println("Start")

	//Multiple defer statements
	//Executes in LIFO order

	defer fmt.Println("END")
	defer add_num(34, 56)
	defer add_num(10, 10)
}
