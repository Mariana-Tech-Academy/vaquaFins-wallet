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

func (h *TransactionHandler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	//request bodyis decoded into transaction struct
	var transaction models.Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	//returns tx from the database
	tx, err := h.Service.GetTransactions(&transaction)
	if err != nil {
		http.Error(w, "transaction not found", http.StatusNotFound)
		return
	}
	//tx is for the transaction from database
	log.Println("Transaction found:", tx)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tx)
}

// GetTransactionsByUserID handles GET requests to retrieve transactions by user ID
func (h *TransactionHandler) GetTransactionsByUserID(w http.ResponseWriter, r *http.Request) {

	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	transactions, err := h.Service.GetTransactionsByUserID(transaction.UserID)
	if err != nil {
		http.Error(w, "transactions not found", http.StatusNotFound)
		return
	}

	log.Println("Transactions found for user ID", transaction.UserID, ":", transactions)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}



/* handler_layer
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
}*/

