package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uint           `gorm:"primarykey" swaggerignore:"true"`
	CreatedAt  time.Time      `swaggerignore:"true"`
	UpdatedAt  time.Time      `swaggerignore:"true"`
	DeletedAt  gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
	FirstName  string         `json:"firstName"`
	LastName   string         `json:"lastName"`
	FatherName *string        `json:"fatherName"`
	Name       string         `json:"username" gorm:"unique"`
	Email      string         `json:"email" gorm:"unique"`
	Password   string         `json:"password"`
	PhotoURL   string         `json:"photo_url"`
}

type AuthUser struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type contextKey string

var (
	ContextUsername = contextKey("username")
)
