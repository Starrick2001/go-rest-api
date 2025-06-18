package main

import (
	"go-rest-api/controllers"
	"go-rest-api/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Account struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	EnvError := godotenv.Load(".env")
	models.Setup()

	if EnvError != nil {
		log.Fatal("Error loading .env file: ", EnvError.Error())
	}
	r := gin.Default()
	api := r.Group("/api/v1")

	api.GET("/", healthCheckHandler)

	auth := api.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	AppPort, isEnvFound := os.LookupEnv("SERVER_PORT")
	if !isEnvFound {
		AppPort = "8080" // Default port if not set in environment variables
		log.Printf("Environment variable APP_PORT not found, using default port %s", AppPort)
	}
	if err := r.Run(":" + AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		panic(err)
	}
}

func healthCheckHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
