package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID           uint   `gorm:"primarykey" json:"id,omitempty"`
	Username     string `json:"username"`
	DisplayName  string `json:"displayName"`
	PasswordHash []byte `json:"-"`
}

func (u *User) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password))
}
