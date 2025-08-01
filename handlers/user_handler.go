package handlers

import "vaqua/service"

// handler layer (handles (http request& reponse) and call the service layer)
//		|
// service layer (business logic and calls the repository layer)
// 		|
// repository layer (handles direct database operations)


type UserHandler struct {
	Service *service.UserService
}