package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Setup() {
	EnvError := godotenv.Load(".env")

	if EnvError != nil {
		log.Fatal("Error loading .env file: ", EnvError.Error())
	}

	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")

	DbUrl := "host=" + DbHost + " port=" + DbPort + " user=" + DbUser + " password=" + DbPassword + " dbname=" + DbName + " sslmode=disable"
	Db, err := gorm.Open(postgres.Open(DbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err.Error())
		panic("Failed to connect to the database")
	}
	log.Println("Connected to the database successfully")
	Db.AutoMigrate(&Account{})
}
