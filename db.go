package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDB() {
	// Connection parameters
	const (
		//host = "localhost" // when not running in container

		host = "db" // when running in container

		port     = 5432
		user     = "user"
		password = "password"
		dbname   = "mydb"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")

	sql := `insert into student values('Lionel Messi','messi@gmail.com');`

	_, err = db.Exec(sql)
	if err != nil {
		log.Fatalf("Error executing query: %v\n", err)
	}
	fmt.Println("Query executed successfully!")
}
