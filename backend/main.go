package main

import (
	"log"
	"net/http"

	"web/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	foo := models.Foo{
		Name: "Hello, World!",
	}

	db, err := gorm.Open(postgres.Open("postgres://postgres:masterpass@postgres-master:5432/postgres"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Foo{})
	db.Create(&foo)

	router.GET("/", func(c *gin.Context) {
		var result []models.Foo
		if err := db.Find(&result).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result, "message": "Data from database"})
	})

	router.Run(":8081")
}
