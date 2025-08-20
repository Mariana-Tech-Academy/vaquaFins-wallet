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
	//transferRepo := &repository.TransferRepo{}
	transactionRepo := &repository.TransactionRepo{}
	//profileRepo := &repository.ProfileRepo{}

	// initialize the service
	userService := &service.UserService{Repo: userRepo}
	//transferService := &service.TransferService{Repo: transferRepo}
	transactionService := &service.TransactionService{Repo: transactionRepo}
	//profileService := &service.ProfileService{Repo: profileRepo}

	// initialize the handler
	userHandler := &handlers.UserHandler{Service: userService}
	//transferHandler := &handlers.TransferHandler{Service: transferService}
	transactionHandler := &handlers.TransactionHandler{Service: transactionService}
	//profileHandler := &handlers.ProfileHandler{Service: profileService}
	healthHandler := &handlers.HealthHandler{}

	// define route
	router := routes.SetupRouter(healthHandler, userHandler,  transactionHandler) //transferHandler,, profileHandler

	fmt.Println("server is running on http://localhost:8080...")
	http.ListenAndServe(":8080", router)
}
