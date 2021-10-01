package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* takes the address of a */
	ptr = &a

	/* takes the address of ptr using address of operator &*/

	pptr = &ptr

	fmt.Printf("Value of a= %d\n", a)
	fmt.Printf("Value available at *ptr= %d\n", *ptr)
	fmt.Printf("Value available at **ptr= %d\n", **pptr)
}
