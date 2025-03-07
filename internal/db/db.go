package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	dbPort := os.Getenv("DATABASE_PORT")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbURL,         // host (this will be "db")
		dbPort,        // port
		"demouser",    // username
		"willnottell", // password
		"tasks",       // database name
	)

	// Try to connect to the database with retries
	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			time.Sleep(time.Second * 2)
			continue
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}

		time.Sleep(time.Second * 2)
	}

	return nil, fmt.Errorf("failed to connect to database after 5 attempts: %v", err)
}
