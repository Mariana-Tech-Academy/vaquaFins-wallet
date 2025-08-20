// handler_layer
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vaqua/models"
	"vaqua/service"
)

type TransactionHandler struct {
	Service *service.TransactionService
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	var transaction models.Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	//call the service_layer
	err = h.Service.CreateTransaction(&transaction)
	if err != nil {
		http.Error(w, "could not create transaction", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	//response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)

}

func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {

	var transaction models.Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	err = h.Service.GetTransactions(&transaction)
	if err != nil {
		http.Error(w, "transaction not found", http.StatusNotFound)
		return
	}
	log.Println("Transaction found:", transaction)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}
