// panic means something went unexpectedly wrong.
package main

import "os"

func main() {
	// we will use panic throughout this site to check for unexpected errors
	//	panic("a problem")

	//A common use of panic is to abort if the function returns an error

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
