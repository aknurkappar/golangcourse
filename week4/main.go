package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
	Age  int64  `json:"not null"`
}

func connectToDatabase() {
	connStr := "host=localhost user=postgres password=secret dbname=postgres port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		fmt.Println("Error in connecting to database:", err)
	}
	fmt.Println("Connected to database!")

	db.AutoMigrate(&User{})
	fmt.Println("Migrations complete")
	createUser(db)
	getUsers(db)
}

func createUser(db *gorm.DB) {
	user := User{Name: "Aknur", Age: 20}
	db.Create(&user)
}

func getUsers(db *gorm.DB) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		fmt.Println("Failed to retrieve users:", err)
	}
	log.Println("Retrieved Users:")
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", u.ID, u.Name, u.Age)
	}
}

func main() {
	fmt.Println("Hello")
	connectToDatabase()
}
