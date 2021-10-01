package main

import (
	"fmt"
	"math"
)

//interface for geometry shapes
type geometry interface {
	area() float64
	perim() float64
}

//for our example we will implement this interface on rect and circle
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

//to implement interface in Go, we just need to implement all the methods in the interface. here we implement geometry on rec

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2*r.height + r.width
}

//implementation for circles

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//If a variable has an interface type, then we can call methods that are in the named interface.
//here's a generic measure function taking advantage of this to work on any geometry.

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 4, height: 3}
	c := circle{radius: 4}
	//circle and rect struct types both implement geometry interface so we can use instances of these structs as arguments to measure
	measure(r)
	measure(c)
}
