package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `json: "username" form:"username"`
	Password string `json: "password" form:"password"`
	Name     string `json: "name" form:"name"`
	Email    string `json: "email" form:"email"`
}
