package databases

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
	"log"
)

type ResponseOrder struct {
	AddressID   uint
	StatusOrder bool
}
type AddressRequest struct {
	Street string
	City   string
	State  string
	Zip    int
}
type OrderDetailRequest struct {
	Address AddressRequest
	Order   ResponseOrder
}

type CartOrder struct {
	ID    uint   `json:"id"`
	Nama  string `json:"nama"`
	Qty   int    `json:"qty"`
	Harga int    `json:"harga"`
}

func CreateOrder(Order_a *models.OrderRequest) (interface{}, error) {

	if err := config.DB.Create(&Order_a.Order).Error; err != nil {
		return nil, err
	}
	return OrderDetailRequest{
		AddressRequest{
			Order_a.Address.Street,
			Order_a.Address.City,
			Order_a.Address.State,
			Order_a.Address.Zip,
		},
		ResponseOrder{
			Order_a.Order.AddressRequest,
			Order_a.Order.StatusOrder,
		},
	}, nil
}

func CreateAddress(address *models.AddressRequest) {
	config.DB.Create(&address)
}

func CreateOrderDet(Order *models.ListOrder) (interface{}, error) {
	if err := config.DB.Create(&Order).Error; err != nil {
		return nil, err
	}
	return Order, nil
}

func GetHargaQtyCart(id int) (int, int, error) {
	cart := models.Cart{}
	err := config.DB.Find(&cart, id)
	if err.Error != nil {
		return 0, 0, err.Error
	}
	log.Println("price", cart.TotalPrice)
	return cart.TotalPrice, cart.Qty, nil
}

func GetOrder(id int) (interface{}, interface{}, error) {
	order := models.ListOrder{}
	type ord struct {
		ID         uint
		Userid     uint
		TotalQty   int
		TotalPrice int
	}

	order_a := models.Order{}

	config.DB.Where(" users_id <> ?", id).Find(&order_a)
	log.Println("id from order", id)
	config.DB.Find(&order)

	return order, ord{order_a.ID, order_a.UsersID, order_a.TotalQty, order_a.TotalPrice}, nil
}
