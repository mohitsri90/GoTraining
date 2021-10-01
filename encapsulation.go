package main

import "fmt"

type Encapsulation struct {
}

func (e *Encapsulation) Expose() {
	fmt.Println("i am exposed")
}

func (e *Encapsulation) hide() {
	fmt.Println("this is super secret")
}

func (e *Encapsulation) Unhide() {
	e.hide()
	fmt.Println("...jk")
}

func main() {
	e := Encapsulation{}
	e.Expose()
	e.Unhide()
}
