package api

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitRoles() {
	adminRole := Role{Name: "admin"}
	userRole := Role{Name: "user"}

	DB.Create(&adminRole)
	DB.Create(&userRole)
}

func Connect() {
	connStr := "host=localhost user=postgres password=secret dbname=postgres port=5433 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in connecting to database:", err)
	}
	fmt.Println("Connected to database!")

	DB.AutoMigrate(&User{}, &Role{})
}
