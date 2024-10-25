package sqlDatabase

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := "host=localhost user=postgres password=secret dbname=postgres port=5433 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error in connecting to database:", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) UNIQUE NOT NULL,
		age INT NOT NULL
	)`)
	if err != nil {
		fmt.Println("Migrations failed:", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS profiles (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		bio VARCHAR(255),
    	profile_picture_url VARCHAR(255),
    	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	)`)
	if err != nil {
		fmt.Println("Migrations failed:", err)
	}

	fmt.Println("Connected to database!")
	return db
}
