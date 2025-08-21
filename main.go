package main

import (
	"github.com/gin-gonic/gin"
	"vaqua/database"
	"vaqua/routes"
)

func main() {
	database.ConnectDatabase()

	// Uncomment once to create demo users
	// database.DB.Create(&models.User;{Name: "Ana", Balance: 100})
	// database.DB.Create(&models.User;{Name: "Aaliyanna", Balance: 50})
	
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}