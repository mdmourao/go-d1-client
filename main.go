package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mdmourao/go-d1"
)

func main() {
	dsn := "http://example.com?token=verysecretpassword&debug=true"
	db, err := sql.Open("god1", dsn)
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer db.Close()

	var n int
	if err := db.QueryRow("SELECT 1").Scan(&n); err != nil {
		log.Fatalf("query: %v", err)
	}

	fmt.Println("SELECT 1 =>", n)
}
