package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wb/internal/service"
	"wb/models"
)

type UserControllerImpl struct {
	userService *service.UserServiceImpl
}

func NewUserControllerImpl(us *service.UserServiceImpl) *UserControllerImpl {
	return &UserControllerImpl{userService: us}
}

func (uc *UserControllerImpl) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := uc.userService.CreateUser(&newUser)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(createdUser)
	http.Error(w, fmt.Sprintf(
		"Всёёёё ок, добавлен пользователь с id: %d и с именем: %s \n", newUser.ID, newUser.Name),
		http.StatusOK)

}
