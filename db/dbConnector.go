package db

import (
	"banplication/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open(""), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Receipt{})
	DB = db
}
