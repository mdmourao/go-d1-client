package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.GET("/mascots", func(c *gin.Context) {
		var mascots []Mascots
		if err := db.Find(&mascots).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, mascots)
	})

	r.PUT("/mascots/:id", func(c *gin.Context) {
		var mascot Mascots
		if err := db.First(&mascot, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Mascot not found"})
			return
		}

		var input struct {
			Description string `json:"description"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mascot.Description = input.Description
		if err := db.Save(&mascot).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, mascot)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server: %v", err)
	}
}
