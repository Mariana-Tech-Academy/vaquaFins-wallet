package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vaqua/models"
	"vaqua/service"
)

func GetTransactionByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	transaction, err := service.GetTransactionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
		return
	}
	created, err := service.CreateTransaction(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Creation failed"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func GetTransactionsByCategory(c *gin.Context) {
	userIDParam := c.Query("user_id")
	category := c.Query("category")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil || category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID or category"})
		return
	}
	txs, err := service.GetTransactionsByCategory(uint(userID), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch transactions"})
		return
	}
	c.JSON(http.StatusOK, txs)
}
func GetIncomeExpenseSummary(c *gin.Context) {
	userIDParam := c.Query("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	income, expense, err := service.GetIncomeExpenseSummary(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not calculate summary"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total_income":  income,
		"total_expense": expense,
	})
}

func GetTransactionsByType(c *gin.Context) {
    userIDParam := c.Query("user_id")
    txType := c.Query("type")
    userID, err := strconv.Atoi(userIDParam)
    if err != nil || txType == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID or type"})
        return
    }
    txs, err := service.GetTransactionsByType(uint(userID), txType)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch transactions"})
        return
    }
    c.JSON(http.StatusOK, txs)
}

func GetTransactionsByUser(c *gin.Context) {
    userIDParam := c.Query("user_id")
    userID, err := strconv.Atoi(userIDParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    txs, err := service.GetTransactionsByUser(uint(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch transactions"})
        return
    }
    c.JSON(http.StatusOK, txs)
}