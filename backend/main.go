package main

import (
	"DuitKito/backend/internal/database"
	"DuitKito/backend/internal/handlers"
	"DuitKito/backend/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// konek ke DB
	database.ConnectDB()
	database.DB.AutoMigrate(&models.Transaction{})

	// routes CRUD
	r.GET("/transactions", handlers.GetTransactions)
	r.POST("/transactions", handlers.CreateTransaction)
	r.PUT("/transactions/:id", handlers.UpdateTransaction)
	r.DELETE("/transactions/:id", handlers.DeleteTransaction)

	r.Run(":8080")
}
