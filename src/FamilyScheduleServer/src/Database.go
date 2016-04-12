package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var (
	name string
)

func main() {
	db, err := sql.Open("postgres", "user=ramik dbname=family sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	
	rows, err := db.Query("SELECT name FROM users")
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name)
	}
	log.Println(err)
}
