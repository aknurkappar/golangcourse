package main

import (
	"fmt"
)

func main() {
	// Exercise 3: Control Structures
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if number == 0 {
		fmt.Println("Entered number is zero: ", number)
	} else if number > 0 {
		fmt.Println("Entered number is positive: ", number)
	} else {
		fmt.Println("Entered number is negative: ", number)
	}

	var sum = 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum of first 10 natural numbers is: ", number)

	var day int
	fmt.Print("Enter a day of week (number): ")
	fmt.Scanln(&day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Please enter valid day of week")
	}
}
