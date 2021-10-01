package main

import "fmt"

const (
	a = iota
	b
	c
	d
)

const (
	a2 = iota
)

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", a2)
}
