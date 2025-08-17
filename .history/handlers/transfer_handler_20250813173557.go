package handlers

import (
	"encoding/json"
	"net/http"
	"vaqua/models"
	"vaqua/service"
)

type TransferHandler struct {
	Service *service.TransferService
}

func (h *TransferHandler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var transfer models.Transfer
	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.Service.CreateTransfer(&transfer)
	if err != nil {
		http.Error(w, "could not complete transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transfer)
}

//Note to self ....i don't think you would want to delete transfer records
// func (h *TransferHandler)DeleteTransfer(w http.ResponseWriter, r *http.Request){
// 	var delTransfer models.Transfer
// 	err := json.NewDecoder(r.Body).Decode(&delTransfer)
// 	if err != nil{
// 		http.Error(w, "invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	err = h.Service.CreateTransfer(&delTransfer)
// 	if err != nil{
// 		http.Error(w, "could not delete transaction", http.StatusInternalServerError)
// 		return
// 	}

// w.WriteHeader(http.StatusCreated)
// json.NewEncoder(w).Encode(delTransfer)
// }
