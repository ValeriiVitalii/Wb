package repositories

import "wb/models"

type IUserRepository interface {
	AddUser(user *models.User) error
	UpdateUser(user models.User) error
	GetUserByID(userID uint) (models.User, error)
}
