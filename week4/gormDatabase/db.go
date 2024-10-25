package gormDatabase

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	connStr := "host=localhost user=postgres password=secret dbname=postgres port=5434 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in connecting to database:", err)
		return nil
	}
	fmt.Println("Connected to database!")

	if err := db.AutoMigrate(&User{}, &Profile{}); err != nil {
		fmt.Println("User migrations failed:", err)
		return nil
	}

	fmt.Println("Migrations complete")
	return db
}
