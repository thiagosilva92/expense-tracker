package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	initDatabase()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/expenses", getExpenses)
	r.POST("/expenses", createExpense)
	r.DELETE("/expenses/:id", deleteExpense)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
