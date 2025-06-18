package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Setup() {
	log.Print("Setting up the database connection...")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")

	DbUrl := "host=" + DbHost + " port=" + DbPort + " user=" + DbUser + " password=" + DbPassword + " dbname=" + DbName + " sslmode=disable"
	var err error
	Db, err = gorm.Open(postgres.Open(DbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err.Error())
		panic("Failed to connect to the database")
	}
	Db.AutoMigrate(&Account{})
	log.Println("Connected to the database successfully")
}
