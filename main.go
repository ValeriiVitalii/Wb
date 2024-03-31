package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"wb/models"
)

func validate(err error, method string) {
	if err != nil {
		fmt.Println("Ошибка validate:", err)
	}
}

func main() {
	dsn := "host=localhost user=postgres password=8823 dbname=wb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	validate(err, "connect")

	// Автомиграция (автоматическое создание таблицы, если она не существует)
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Ошибка при создании таблицы:", err)
	}

	user := models.User{Name: "john_doe"}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal("Ошибка при создании пользователя:", result.Error)
	}
	log.Println("Создан пользователь:", user)
	//Добавить метод validate у которого будет в аргументах err + название метода из которого вызывается

}
