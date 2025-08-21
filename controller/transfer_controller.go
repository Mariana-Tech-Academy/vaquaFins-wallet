package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vaqua/models"
	"vaqua/service"
)

func TransferMoney(c *gin.Context) {
	var req models.TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	if err := service.TransferMoney(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transfer successful!"})
}
