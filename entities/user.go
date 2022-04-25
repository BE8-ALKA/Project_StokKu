package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"author"`
	Password string `json:"password" form:"password"`
}
