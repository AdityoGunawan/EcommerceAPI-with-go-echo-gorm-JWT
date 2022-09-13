package models

import "gorm.io/gorm"

// struktur data product
type Product struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Category    string `json:"category" form:"category"`
	Description string `json:"description" form:"description"`
	UsersID     uint
	Cart        []Cart
}
