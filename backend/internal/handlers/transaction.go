package handlers

import (
	"DuitKito/backend/internal/database"
	"DuitKito/backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	database.DB.Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c *gin.Context) {
	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

func UpdateTransaction(c *gin.Context) {
	var t models.Transaction
	id := c.Param("id")

	if err := database.DB.First(&t, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t.Description = input.Description
	t.Amount = input.Amount
	t.Type = input.Type

	database.DB.Save(&t)
	c.JSON(http.StatusOK, t)
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	var t models.Transaction

	if err := database.DB.First(&t, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	database.DB.Delete(&t)
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted"})
}
