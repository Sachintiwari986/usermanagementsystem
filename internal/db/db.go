package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

var DB *sql.DB

func InitDB() {
	var err error
	// Use the service name 'db' as the host if using Docker
	DB, err = sql.Open("mysql", "user:password@tcp(db:3306)/user_management")
	if err != nil {
		log.Fatal(err)
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        reset_token VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the table exists
	checkTableQuery := `SELECT 1 FROM users LIMIT 1;`
	_, err = DB.Exec(checkTableQuery)
	if err != nil {
		log.Println("Table 'users' does not exist or cannot be accessed.")
	} else {
		log.Println("Table 'users' exists and is accessible.")
	}
}
