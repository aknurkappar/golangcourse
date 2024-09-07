package main

import (
	"fmt"
)

func main() {

	// Exercise 1: Hello, World!
	fmt.Println("Hello Go!")

	// Exercise 2: Variables and Data Types
	var i int = 42
	i2 := 42
	fmt.Println("Integer", i, i2)

	var f float64 = float64(i)
	f2 := float64(i)
	fmt.Println("Float", f, f2)

	var u uint = uint(f)
	u2 := uint(f)
	fmt.Println("Non-negative integers", u, u2)

	var myName = "Aknur"
	myName2 := "Aknur"
	fmt.Println("String", myName, myName2)

	var isTrue bool = true
	isFalse := false
	fmt.Println("Boolean", isTrue, isFalse)
}
