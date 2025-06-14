package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Account struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Handler struct {
	db *gorm.DB
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db: " + err.Error())
	}
	db.AutoMigrate(&Account{})
	handler := newHandler(db)

	r := gin.Default()
	r.GET("/", handler.healthCheckHandler)
	r.GET("/accounts", handler.listAccountHandler)
	r.POST("/accounts", handler.createAccountHandler)
	r.DELETE("/accounts/:id", handler.deleteAccountHandler)

	r.Run()
}

func (handler *Handler) healthCheckHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
func (handler *Handler) listAccountHandler(ctx *gin.Context) {
	var accounts []Account
	if result := handler.db.Find(&accounts); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &accounts)
}
func (handler *Handler) createAccountHandler(ctx *gin.Context) {
	var account Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if account.ID == "" || account.Email == "" || account.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account"})
		return
	}
	if result := handler.db.Create(&account); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Ok"})
}
func (handler *Handler) deleteAccountHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if result := handler.db.Delete(&Account{}, id); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}
