package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID           uint   `json:"id,omitempty" gorm:"primarykey"`
	Username     string `json:"username" gorm:"unique"`
	DisplayName  string `json:"displayName"`
	PasswordHash []byte `json:"-"`

	Blog Blog
}

func (u *User) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password))
}
