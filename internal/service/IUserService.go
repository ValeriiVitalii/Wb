package service

import "wb/models"

type IUserService interface {
	CreateUser(user *models.User) error
	UpdateUser(user models.User) error
	GetUserByID(user models.User) models.User
}
