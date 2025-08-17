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



func(s *TransferHandler) TransferMoney(w http.ResponseWriter, r *http.Request){

var accs models.Transfer
err := json.NewDecoder(r.Body).Decode(&accs)
if err != nil {
	http.Error(w, "unable to decode",http.StatusBadRequest)
	return
}

err = s.Service.TransferMoney(accs.AccountNum,accs.RecipientID,accs.Amount)
if err != nil{
	http.Error(w, "unable to transfer money", http.StatusBadGateway)
	return
}

//response
w.Write([]byte("transfer successful"))

}