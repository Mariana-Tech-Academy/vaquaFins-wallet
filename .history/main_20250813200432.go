package main

import (
	"fmt"
	"net/http"
	"vaqua/config"
	"vaqua/db"
	"vaqua/handlers"
	"vaqua/repository"
	"vaqua/routes"
	"vaqua/service"
)

func main() {

	// Loading up variables
	config.LoadEnv()

	// connect to database
	db.InitDb()

	// initialize the repo
	userRepo := &repository.UserRepo{}
	// transferRepo := &repository.TransferRepo{}
	transactionRepo := &repository.TransactionRepo{}

	// initialize the service
	userService := &service.UserService{Repo: userRepo}
	// transferService := &service.TransferService{Repo: transferRepo}
	transactionService := &service.TransactionService{Repo: transactionRepo}

	// initialize the handler
	userHandler := &handlers.UserHandler{Service: userService}
	// transferHandler := &handlers.TransferHandler{Service: transferService}
	transactionHandler := &handlers.TransactionHandler{Service: transactionService}

	healthHandler := &handlers.HealthHandler{}

	// define route

	router := routes.SetupRouter(healthHandler, userHandler, transactionHandler) //, transferHandler <--include after testing

	fmt.Println("server is running on localhost:8080...")
	http.ListenAndServe(":8080", router)
}
