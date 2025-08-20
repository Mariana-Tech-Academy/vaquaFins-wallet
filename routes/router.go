package routes

import (
	// "net/http"
	"vaqua/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter(healthCheckHandler *handlers.HealthHandler,
	userHandler *handlers.UserHandler,
	//transferHandler *handlers.TransferHandler,
	//profileHandler *handlers.ProfileHandler,
	transactionHandler *handlers.TransactionHandler,) *mux.Router {
	
	/* can delete?
	   userHandler *handlers.UserHandler, transactionHandler *handlers.TransactionHandler, transferHandler *handlers.TransferHandler
	*/

	r := mux.NewRouter()

	//public routes
	r.HandleFunc("/healthCheck", healthCheckHandler.HealthCheck)
	r.HandleFunc("/register", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LogIn).Methods("POST")
	r.HandleFunc("/transaction", transactionHandler.CreateTransaction).Methods("POST")
	r.HandleFunc("/transaction", transactionHandler.GetTransactions).Methods("GET")
	// r.HandleFunc("/transfer", transferHandler.TransferMoney).Methods("POST")
	// r.HandleFunc("/profile", profileHandler.UpdateProfile).Methods("PUT")
	// r.HandleFunc("/profile", profileHandler.GetProfile).Methods("GET")
	// r.HandleFunc("/profile", profileHandler.DeleteProfile).Methods("DELETE")

	//this is for retrieving all Transactions
	// r.HandleFunc("/transaction", userHandler.Transaction).Methods("Get")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use()

	//Editor-only routes
	//protected.Handle()
	//protected.Handle()

	//Authenticated user routes
	//protected.HandleFunc()
	//protected.HandleFunc()
	return r

}
