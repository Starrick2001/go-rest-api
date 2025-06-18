package models

import (
	"go-rest-api/utils"
	"html"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	// Account represents a user account in the system.
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

func (a *Account) SaveAccount() (*Account, error) {
	// SaveAccount saves the account to the database.
	if err := Db.Create(&a).Error; err != nil {
		log.Print(err.Error())
		return &Account{}, err
	}
	return a, nil
}

func (a *Account) BeforeSave(tx *gorm.DB) (err error) {
	// BeforeSave is a GORM hook that runs before saving the account.
	if a.Email == "" {
		return gorm.ErrInvalidData
	}
	if a.Password == "" {
		return gorm.ErrInvalidData
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hashedPassword)
	a.Email = html.EscapeString(strings.TrimSpace(a.Email))
	return nil
}

func VerifyPassword(providedPassword, storedPassword string) error {
	// VerifyPassword checks if the provided password matches the stored password.
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
}
func CheckLogin(email string, password string) (string, error) {
	account := &Account{}
	if err := Db.Where("email = ?", email).First(&account).Error; err != nil {
		log.Printf("Error finding account: %v", err)
		return "", err
	}

	if err := VerifyPassword(password, account.Password); err != nil {
		log.Printf("Invalid password for account with email %s: %v", email, err)
		return "", err
	}

	token, err := utils.GenerateToken(account.ID)
	if err != nil {
		log.Printf("Error generating token for account with email %s: %v", email, err)
		return "", err
	}
	return token, nil
}
