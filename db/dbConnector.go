package db

import (
	"banplication/model"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var _ *gorm.DB

func loadCredentials() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("Successfully loaded .env file")
}

func Connect() {
	loadCredentials()
	host := os.Getenv("HOST_NAME")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("user")
	dbname := os.Getenv("dbname")
	passwd := os.Getenv("DB_PASSWD")
	sslMode := os.Getenv("SSL_MODE")
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, passwd, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	aMigrationErr := db.AutoMigrate(&model.Receipt{})
	if aMigrationErr != nil {
		panic(aMigrationErr)
	}

	_ = db
}
