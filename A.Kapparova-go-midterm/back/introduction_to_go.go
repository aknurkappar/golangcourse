package main

import (
	"fmt"
)

type Task struct {
	title  string
	status string
}

func NewTask(title string) Task {
	return Task{
		title:  title,
		status: "todo",
	}
}

func (task *Task) displayTask() string {
	return task.status
}

func (task *Task) changeStatus(status string) {
	task.status = status
	fmt.Printf("%s: status changed to %s!", task.title, task.status)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func RunIntroductionToGo() {
	var x, y int

	fmt.Print("Enter first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter second number: ")
	fmt.Scan(&y)

	sum := add(x, y)
	difference := subtract(x, y)

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", difference)

	task := NewTask("My Task")
	task.changeStatus("done")
}
