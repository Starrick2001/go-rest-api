package controllers

import (
	"go-rest-api/models"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func Register(ctx *gin.Context) {
	// This function will handle user registration
	var body RegisterRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}
	account := models.Account{
		Email:    body.Email,
		Password: body.Password, // In a real application, you should hash the password before saving it
	}
	if _, err := account.SaveAccount(); err != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to save account: " + err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "User registered successfully",
	})
}
