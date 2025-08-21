package routes

import (
	// "net/http"
	"vaqua/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter(healthCheckHandler *handlers.HealthHandler,
	userHandler *handlers.UserHandler,

	transactionHandler *handlers.TransactionHandler) *mux.Router {
	r := mux.NewRouter()

	//public routes
	r.HandleFunc("/healthCheck", healthCheckHandler.HealthCheck)
	r.HandleFunc("/register", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LogIn).Methods("POST")
	r.HandleFunc("/transaction", transactionHandler.GetTransaction).Methods("GET")
	r.HandleFunc("/transactions", transactionHandler.GetTransactionsByUserID).Methods("GET")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use()
	return r

}
