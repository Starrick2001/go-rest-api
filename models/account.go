package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	// Account represents a user account in the system.
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

func (a *Account) SaveAccount() (*Account, error) {
	// SaveAccount saves the account to the database.
	if err := Db.Create(a).Error; err != nil {
		return &Account{}, err
	}
	return a, nil
}
