package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func connectToDatabase() {
	connStr := "user=postgres password=secret dbname=postgres port=5433 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		fmt.Println("Error in connecting to database:", err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println("Error in pinging database:", err)
	}
}

func main() {
	fmt.Println("Hello")
	connectToDatabase()
}
