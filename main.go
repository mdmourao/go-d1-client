package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mdmourao/go-d1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("no .env file loaded: %v", err)
	}

	dsn := os.Getenv("D1_DSN")
	if dsn == "" {
		log.Fatal("D1_DSN is not set")
	}

	sqlDB, err := sql.Open("god1", dsn)
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer sqlDB.Close()

	db, err := gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm open: %v", err)
	}

	var mascots []Mascots
	if err := db.Find(&mascots).Error; err != nil {
		log.Fatalf("find: %v", err)
	}

	for i, m := range mascots {
		fmt.Printf("%d - %d %s %s — %s\n", i, m.ID, m.Language, m.Name, m.Description)
	}
}
