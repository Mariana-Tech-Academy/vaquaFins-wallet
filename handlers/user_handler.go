package handlers

import (
	"encoding/json"
	"net/http"
	"vaqua/models"
	"vaqua/service"
)

type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// calls the service layer
	err = h.Service.CreateUser(&user)
	if err != nil {
		http.Error(w, "could not register user", http.StatusInternalServerError)
		return
	}
	err = h.Service.CreateUser(&user)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		http.Error(w, "user already exist", http.StatusAccepted)
	}

	//response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	var request models.User

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	token, err := h.Service.LogInUser(request)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusInternalServerError)
		return
	}

	//response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)

}
