package models

import "gorm.io/gorm"

// struktur data cart
type Cart struct {
	gorm.Model
	Qty        int  `json:"qty" form:"qty"`
	TotalPrice int  `json:"total_harga" form:"total_price"`
	ProductID  uint `json:"product_id" form:"product_id"`
	UsersID    uint
	ListOrder  []ListOrder
}
