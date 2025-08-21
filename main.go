package main

import (
	"log"
	"net/http"
	"vaqua/config"
	"vaqua/db"
	"vaqua/handlers"
	"vaqua/repository"
	"vaqua/routes"
	"vaqua/service"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	db.InitDb()

	// Initialize repositories
	userRepo := &repository.UserRepo{}
	transactionRepo := &repository.TransactionRepo{}

	// Initialize services
	userService := &service.UserService{Repo: userRepo}
	transactionService := &service.TransactionService{Repo: transactionRepo}

	// Initialize handlers
	userHandler := &handlers.UserHandler{Service: userService}
	transactionHandler := &handlers.TransactionHandler{Service: transactionService}
	healthHandler := &handlers.HealthHandler{}

	// Setup routes
	router := routes.SetupRouter(healthHandler, userHandler, transactionHandler)

	log.Println("Server is running on http://localhost:8080...")
	http.ListenAndServe(":8080", router)
}
