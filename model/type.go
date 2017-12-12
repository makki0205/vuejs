package model

import "github.com/jinzhu/gorm"

type Model struct {
	gorm.Model
}
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
