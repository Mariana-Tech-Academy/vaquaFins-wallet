package routes

import (
	// "net/http"
	"vaqua/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter(healthCheckHandler *handlers.HealthHandler,
	userHandler *handlers.UserHandler,
	transferHandler *handlers.TransferHandler,
	transactionHandler *handlers.TransactionHandler,
	incomeAndExpensesHandler *handlers.IncomeAndExpensesHandler) *mux.Router {

	r := mux.NewRouter()

	//public routes
	r.HandleFunc("/healthCheck", healthCheckHandler.HealthCheck)
	r.HandleFunc("/register", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LogIn).Methods("POST")
	r.HandleFunc("/transaction", transactionHandler.GetTransaction).Methods("GET")
	r.HandleFunc("/transactions", transactionHandler.GetTransactionsByUserID).Methods("GET")
	r.HandleFunc("/transfer", transferHandler.TransferMoney).Methods("POST")
    r.HandleFunc("/accounts/{id:[0-9]+}/summary", incomeAndExpensesHandler.GetSummary).Methods("GET")

	//this is for retrieving all Transactions
	// r.HandleFunc("/transaction", userHandler.Transaction).Methods("Get")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use()

	//Editor-only routes
	//protected.Handle()
	//protected.Handle()

	return r

}
