package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vaqua/models"
	"vaqua/service"
)

// handler layer (handles (http request& response) and call the service layer)
//		|
// service layer (business logic and calls the repository layer)
// 		|
// repository layer (handles direct database operations)

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
	// call the service layer
	err = h.Service.CreateUser(&user)
	if err != nil {
		http.Error(w, "could not register user", http.StatusInternalServerError)
		return
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

/*func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
    tokenString := r.Header.Get("Authorization") //Get the token from the request header

    //Check if the token is present
    if tokenString == "" {
        http.Error(w, "authorization token is missing", http.StatusUnauthorized)
        return
    }

    //Pass the token to the service layer for invalidation
    err := h.Service.LogoutUser(tokenString)
    if err != nil {
        http.Error(w, "failed to logout", http.StatusInternalServerError)
        return
    }

    //Response
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "logout successful"})
}*/

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
    tokenString := r.Header.Get("Authorization")

    if tokenString == "" {
        http.Error(w, "authorization token is missing", http.StatusUnauthorized)
        return
    }

    // Remove "Bearer " prefix if present
    if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
        tokenString = tokenString[7:]
    }

    err := h.Service.LogoutUser(tokenString)
    if err != nil {
        // Log the real error
        fmt.Println("Logout error:", err)
        http.Error(w, "failed to logout", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "logout successful"})
}

