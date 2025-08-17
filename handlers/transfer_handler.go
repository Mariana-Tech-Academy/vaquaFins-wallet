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




//4)Returns either an error slip (4xx/5xx) or a success receipt with transfer_id.

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