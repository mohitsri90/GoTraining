//Go suports pointers, allowing you to pass references.
package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

//zeroptr has an *int parameter, meaning it takes int pointer
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial: ", i)
	zeroVal(i)
	fmt.Println("zeroVal: ", i)
	//The &i gives memory address of i, i.e pointer to i
	zeroptr(&i)
	fmt.Println("zeroptr: ", i)
	fmt.Println("zeroptr: ", &i)
}
