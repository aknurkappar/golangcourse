package main

import (
	"main/api"
)

func main() {
	api.Connect()
	api.Run()

	//api.CreateCategory("Study")
	//api.CreateCategory("Personal")
	//
	//api.CreateTask("Do midterm project", "todo", 1)
	//api.CreateTask("Write report", "todo", 1)
	//tasks := api.GetTasks()
	//fmt.Println("All tasks:", tasks)
	//
	//studyTasks, err := api.GetTasksByCategory(1)
	//if err != nil {
	//	fmt.Println("Error retrieving tasks:", err)
	//	return
	//}
	//fmt.Println("Study tasks:", studyTasks)
	//
	//api.UpdateTask(1, "Do midterm project", "in progress")
	//api.DeleteTask(1)
}
