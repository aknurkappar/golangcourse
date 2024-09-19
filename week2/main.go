package main

import (
	"fmt"
)

type BasicUser interface {
	signIn()
}
type Department struct {
	Id   int
	Name string
}
type User struct {
	Id   string
	Name string
}
type Student struct {
	User
	Year int
	Department
}
type Teacher struct {
	User
	Degree string
	Department
}

func (student Student) signIn() {
	fmt.Println("Student of ", student.Year, "year", student.User.Name, "signed in! Show student's interface")
}

func (teacher Teacher) signIn() {
	fmt.Println("Teacher", teacher.User.Name, "(", teacher.Degree, ") signed in! Show teacher's interface")
}

func (teacher Teacher) getDegree() string {
	return teacher.Degree
}

func (teacher *Teacher) changeDegree(newDegree string) {
	teacher.Degree = newDegree
}

func main() {
	var itDepartment = Department{
		Id:   1,
		Name: "SITE",
	}
	student := &Student{
		User: User{
			Id:   "21B03",
			Name: "Aknur",
		},
		Year:       4,
		Department: itDepartment,
	}
	teacher := &Teacher{
		User: User{
			Id:   "2000",
			Name: "Azamat",
		},
		Degree:     "PHD",
		Department: itDepartment,
	}

	fmt.Println(itDepartment)
	fmt.Println(student)
	fmt.Println(teacher)
	student.signIn()
	teacher.signIn()

	fmt.Println(teacher.getDegree())
	teacher.changeDegree("Doctor")
	fmt.Println(teacher.getDegree())
}
