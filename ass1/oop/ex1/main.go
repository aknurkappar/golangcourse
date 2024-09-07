package main

import (
	"fmt"
)

type Person struct {
	name string
	age  uint
}

func (p Person) greet() {
	fmt.Printf("Hello, I'm %s!", p.name)
}

func main() {
	p1 := Person{name: "Aknur", age: 20}
	p1.greet()
}
