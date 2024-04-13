package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"wb/internal/controllers"
	"wb/internal/repositories"
	"wb/internal/service"
	"wb/models"
)

func validate(err error, method string) {
	if err != nil {
		if method == "connect" {
			log.Fatal("Ошибка коннекта к базе:", err)
		}
		if method == "create_table_users" {
			log.Fatal("Ошибка при создании таблицы:", err)
		}
		if method == "create_user" {
			log.Fatal("Ошибка при создании пользователя:", err)
		}
	}
}

func main() {
	dsn := "host=localhost user=postgres password=8823 dbname=wb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	validate(err, "connect")

	userService := service.NewUserServiceImpl(repositories.NewUserRepositoryImpl(db))

	userService.CreateUser(&models.User{Name: "Elуууу"})
	user := userService.GetUserByID(12)

	fmt.Println(user.Name)

	r := mux.NewRouter()
	userController := controllers.NewUserControllerImpl(userService)

	r.HandleFunc("/users", userController.CreateUserHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}
