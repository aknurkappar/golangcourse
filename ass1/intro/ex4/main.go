package main

import (
	"fmt"
)

func add(param1, param2 int) int {
	return param1 + param2
}

func swap(a, b string) (string, string) {
	return b, a
}

func quotientAndRemainder(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	// Exercise 4: Functions and Multiple Return Values
	fmt.Println(add(34, 45))

	// swap()
	var a, b string
	fmt.Print("Enter 2 strings: ")
	fmt.Scanln(&a, &b)

	aSwapped, bSwapped := swap(a, b)
	fmt.Println("In reversed order: ", aSwapped, bSwapped)

	// quotientAndRemainder
	var num1, num2 int
	fmt.Print("Enter 2 numbers: ")
	fmt.Scanln(&num1, &num2)

	quotient, remainder := quotientAndRemainder(num1, num2)
	fmt.Printf(
		"Quotient: %d, remainder: %d of %d/%d \n",
		quotient, remainder, num1, num2)
}
