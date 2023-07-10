package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model

	UserID uint `json:"userId"`

	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
