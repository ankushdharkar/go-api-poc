package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context) {
		typeVal := c.DefaultQuery("type", "no type found")
		isValid := c.Query("isValid")

		c.JSON(200, gin.H{
			"type":    typeVal,
			"isValid": isValid,
			"message": "pong",
		})
	})

	router.Run(":" + port)
}
