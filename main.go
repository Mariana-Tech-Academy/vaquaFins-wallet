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
	incomeAndExpensesRepo := &repository.IncomeAndExpensesRepo{}
	transferRepo := &repository.TransferRepo{}

	// initialize the service
	userService := &service.UserService{Repo: userRepo}
	// transferService := &service.TransferService{Repo: transferRepo}
	transactionService := &service.TransactionService{Repo: transactionRepo}
	incomeAndExpensesService:=&service.IncomeAndExpensesService{Repo: incomeAndExpensesRepo}
	transferService:= &service.TransferService{Repo: transferRepo}

	// initialize the handler
	userHandler := &handlers.UserHandler{Service: userService}
	// transferHandler := &handlers.TransferHandler{Service: transferService}
	transactionHandler := &handlers.TransactionHandler{Service: transactionService}
	incomeAndExpensesHandler := &handlers.IncomeAndExpensesHandler{Service: incomeAndExpensesService}
	healthHandler := &handlers.HealthHandler{}
	transferHandler:= &handlers.TransferHandler{Service: transferService}

	// define route


	
	router := routes.SetupRouter(healthHandler, userHandler,transferHandler,transactionHandler,incomeAndExpensesHandler,) //, transferHandler <--include after testing

	fmt.Println("server is running on localhost:8080...")
	http.ListenAndServe(":8080", router)
}
