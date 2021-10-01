package main

import "fmt"

//this function intSeq returns another function, which we anomously define in the body
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// we call intSeq, assigning the result(a function) to nextInt
	nextInt := intSeq()
	// see the effect of closure by calling it few times
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
