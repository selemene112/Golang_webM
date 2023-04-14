package database

import (
	"fmt"
	"log"

	"final/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	DB_HOST     = "localhost"
	DB_USER     = "admin"
	DB_PASSWORD = "lele123"
	DB_NAME     = "final"
	DB_PORT     = 5432
	
	db *gorm.DB
	err error
)

func StartDB() {

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	dsn := config 
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Connection To DataBase : ", err)
	}
	db.AutoMigrate(&models.Comment{},&models.SosialMedias{})

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{})//, models.Admin{}

	
}

func GetDB() *gorm.DB {
	return db
}



