package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//Suppose we have function call f(s). Here's how we would call that
	f("direct")
	//to invoke this in go routine use go f(s)
	go f("goroutine")
	// you can also start go routine for anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//our two function calls are running asynchronously in separate goroutines now
	fmt.Scanln()
	fmt.Println("done")
}
