package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Printf("%T", d)
	fmt.Println(int64(d))
	fmt.Printf("%T", d)
	fmt.Println(math.Sin(n))
}
