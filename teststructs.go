//Go's structs are typed collection of fields. They are useful for grouping data together to form records.
package main

import "fmt"

//This person struct type has name and age fields
type person struct {
	name string
	age  int
}

func main() {
	//this syntax creates new struct
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Mohit", age: 31})
	//Omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})
	//&prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})
	//Access struct fields with a dot .
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	//you can also use dot with struct pointers - the pointers are automatically deferenced
	sp := &s
	fmt.Println(sp.age)
	//Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
