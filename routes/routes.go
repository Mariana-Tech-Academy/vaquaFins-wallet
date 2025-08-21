package routes

import (
	"github.com/gin-gonic/gin"
	"vaqua/controller"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/transactions/:id", controller.GetTransactionByID)
    router.POST("/transactions", controller.CreateTransaction)
    router.PATCH("/transfer", controller.TransferMoney)
    router.GET("/transactions", controller.GetTransactionsByUser)
    router.GET("/transactions/type", controller.GetTransactionsByType)
    router.GET("/transactions/category", controller.GetTransactionsByCategory)
    router.GET("/transactions/summary", controller.GetIncomeExpenseSummary)
}
