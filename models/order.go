package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	StatusOrder    bool `json:"status_order" form:"status_order"`
	TotalQty       int  `json:"total_qty" form:"total_qty"`
	TotalPrice     int  `json:"total_price" form:"total_price"`
	AddressRequest uint
	UsersID        uint
	ListOrder      []ListOrder
}

type ListOrder struct {
	gorm.Model
	OrderID uint
	CartID  uint
}

type AddressRequest struct {
	gorm.Model
	Street string `json:"street" form:"street"`
	City   string `json:"city" form:"city"`
	State  string `json:"state" form:"state"`
	Zip    int    `json:"zip" form:"zip"`
}

type OrderRequest struct {
	DetailCartId []int          `json:"detail_cart_id" form:"detail_cart_id"`
	Order        Order          `json:"order" form:"order"`
	Address      AddressRequest `json:"address" form:"address"`
}
