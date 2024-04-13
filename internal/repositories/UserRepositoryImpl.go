package repositories

import (
	"gorm.io/gorm"
	"wb/models"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) AddUser(user *models.User) error {
	result := ur.db.Create(&user)
	return result.Error
}

func (ur *UserRepositoryImpl) UpdateUser(user models.User) error {
	result := ur.db.Update("Name", &user)
	return result.Error
}

func (ur *UserRepositoryImpl) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
