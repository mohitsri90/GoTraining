package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	// if you want only subset of returned values use blank
	_, c := vals()
	fmt.Println(c)
}
