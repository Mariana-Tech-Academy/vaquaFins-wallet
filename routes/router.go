package routes

import (
	// "net/http"
	"vaqua/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter(healthCheckHandler *handlers.HealthHandler,
	userHandler *handlers.UserHandler,
	transferHandler *handlers.TransferHandler,
	transactionHandler *handlers.TransactionHandler) *mux.Router {
	/* can delete?
	   userHandler *handlers.UserHandler, transactionHandler *handlers.TransactionHandler, transferHandler *handlers.TransferHandler
	*/

	r := mux.NewRouter()

	//public routes
	r.HandleFunc("/healthCheck", healthCheckHandler.HealthCheck)
	r.HandleFunc("/signup", userHandler.SignUp).Methods("POST")
	r.HandleFunc("/login", userHandler.LogIn).Methods("POST")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use()

	//Editor-only routes
	//protected.Handle()
	//protected.Handle()

	//Authenticed user routes
	//protected.HandleFunc()
	//protected.HandleFunc()
	return r

}
