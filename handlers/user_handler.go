package handlers

import (
	"encoding/json"
	"net/http"
	"vaqua/models"
	"vaqua/service"
)

// handler layer (handles (http request& reponse) and call the service layer)
//		|
// service layer (business logic and calls the repository layer)
// 		|
// repository layer (handles direct database operations)


type UserHandler struct {
	Service *service.UserService
}


func (u *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "invalid request body", http.StatusBadRequest)
        return
    }
    // call the service layer
    err = u.Service.RegisterUser(&user)
    if err != nil {
        http.Error(w, "could not register user", http.StatusInternalServerError)
        return
    }
    //response
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var request models.User

    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, "invalid request body", http.StatusBadRequest)
        return
    }
    token, err := u.Service.LoginUser(request)
    if err != nil {
        http.Error(w, "invalid credentials", http.StatusInternalServerError)
        return
    }
    //response
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(token)

}


