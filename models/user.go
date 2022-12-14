package models

import "gorm.io/gorm"

// struktur data users
type Users struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string
	Product  []Product
	Cart     []Cart
	Order    []Order
}
