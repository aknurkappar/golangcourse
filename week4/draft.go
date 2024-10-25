package main

import (
	"encoding/json"
	"fmt"
)

func (user User) toJson() string {
	data, err := json.Marshal(user)

	if err != nil {
		fmt.Println("Error in converting to json:", err)
		return ""
	}
	return string(data)
}

func main() {
	fmt.Println("Hello")
	user := User{ID: 12, Name: "Aknur", Age: 20}
	var jsonData = user.toJson()
	fmt.Println("json format:", jsonData)
	connectToDatabase()
}
