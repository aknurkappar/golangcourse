package main

import (
	"fmt"
)

type Employee struct {
	id   uint
	name string
}

type Manager struct {
	Employee
	department string
}

func (e Employee) work() {
	fmt.Printf("Employee %s with id %d", e.name, e.id)
}

func main() {
	e1 := Employee{name: "Aknur", id: 123}
	e1.work()

	m1 := Manager{
		Employee: Employee{
			name: "Aruna", id: 145,
		},
		department: "Development",
	}
	m1.Employee.work()

	fmt.Println(m1)
}
