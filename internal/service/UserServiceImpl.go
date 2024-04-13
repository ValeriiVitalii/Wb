package service

import (
	"errors"
	"fmt"
	"log"
	"wb/internal/repositories"
	"wb/models"
)

type UserServiceImpl struct {
	userRepository *repositories.UserRepositoryImpl
}

func NewUserServiceImpl(ur *repositories.UserRepositoryImpl) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: ur,
	}
}

func (us *UserServiceImpl) CreateUser(user *models.User) error {
	validate(user.Name)

	err := us.userRepository.AddUser(user)

	checkError(err, "Ошибка при создании пользователя:")

	log.Println(fmt.Sprintf("Добавлен user с id: %d и с именем: %s\n", user.ID, user.Name))
	return err
}

func (us *UserServiceImpl) UpdateUser(user models.User) error {
	err := us.userRepository.UpdateUser(user)

	log.Println(fmt.Sprintf("Обновлено имя user на: %s\n", user.Name))
	return err
}

func (us *UserServiceImpl) GetUserByID(userID uint) *models.User {
	user, err := us.userRepository.GetUserByID(userID)

	if err != nil {
		log.Fatal("Ошибка нахождения пользователя")
	}

	log.Println(fmt.Sprintf("Найден пользоваль: %s\n", user.Name))
	return user
}

func validate(name string) {
	if len(name) < 3 {
		log.Fatal(errors.New("имя пользователя должно содержать минимум 3 буквы"))
	}
}

func checkError(err error, errorMessage string) {
	if err != nil {
		log.Fatal(errorMessage, err)
	}
}
