package main

import (
	"fmt"
	"reflect"
)

//Name of struct tag used in example
const tagName = "Validate"

type User struct {
	Id    int    `validate:"="`
	Name  string `validate:"presence,min=2,max=32"`
	Email string `validate:"email,required"`
}

func main() {
	user := User{
		Id:    1,
		Name:  "Mohit Srivastav",
		Email: "mohit@example.com",
	}

	//typeof returns the reflection type that represents the dynamic type of variable.
	//if a variable is nil interface value, TypeOf returns nil

	t := reflect.TypeOf(user)

	//Get the type and kind of user variable
	fmt.Println("Type: ", t.Name())
	fmt.Println("Kind: ", t.Kind())

	//Iterate over all available fields and read the tag values
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		//get the field tag value
		tag := field.Tag.Get(tagName)

		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}
