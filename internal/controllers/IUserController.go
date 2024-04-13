package controllers

import "net/http"

type IUserController interface {
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
}
