package calculator

import "fmt"

func Calc() {
	fmt.Println("Welcome to Calculator")
	fmt.Println("********MAIN MENU***********")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("****************************")
	var choice int
	fmt.Println("Please select one of the option from the menu")
	//choice will store the users input from menu shown above
	fmt.Scan(&choice)
	var a, b int

	//After the choice of operation, user will be asked to enter 2 int values one by one
	// to perform the operation on
	fmt.Println("Enter the value of a: ")
	fmt.Scan(&a)
	fmt.Println("Enter the value of b: ")
	fmt.Scan(&b)

	if choice == 1 {
		ans := a + b
		fmt.Println("Answer= ", ans)
	} else if choice == 2 {
		ans := a - b
		fmt.Println("Answer= ", ans)
	} else if choice == 3 {
		ans := a * b
		fmt.Println("Answer= ", ans)
	} else {
		ans := a / b
		fmt.Println("Answer= ", ans)
	}
	fmt.Println("********Thank you for using the calculator********")
}
