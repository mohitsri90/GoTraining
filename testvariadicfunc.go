package main

import "fmt"

// this function takes arbitrary number of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	//variadic functions can be called in ususL WAY with individual args
	sum(1, 2)
	sum(1, 2, 3)
	// if you already have multiple args in a slice, apply them to a variadic function
	nums := []int{1, 2, 3, 4, 1}
	sum(nums...)
}
