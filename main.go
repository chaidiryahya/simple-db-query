package main

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "user"
	PASSWORD = "pass"
	DBNAME   = "books"
)

func main() {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME,
	)

	DB, err := sql.Open("postgres", connString)
	if err != nil {
		log.Println("Fail DB Connection")
		log.Fatal(err)
	}
	defer DB.Close()
	log.Println("DB Connected")

	id := 1
	query := "DELETE FROM books WHERE id = $1"
	_, err = DB.Query(query, id)
	if err != nil {
		log.Println("failed to execute query", err)
		return
	}
}
