package databases

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
	"fmt"
)

func GetAllCart(id_user_token int) (interface{}, error) {
	type result struct {
		ID          uint   `json:"id"`
		Qty         int    `json:"qty"`
		TotalPrice  int    `json:"total_price"`
		UsersID     uint   `json:"users_id"`
		ProductID   uint   `json:"product_id"`
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Category    string `json:"category"`
		Description string `json:"description"`
	}
	cart := []result{}
	where_clause := fmt.Sprintf("carts.users_id = %v", id_user_token)

	// query join untuk menampilkan struktur data cart
	query := config.DB.Table("carts").Select("carts.id, carts.qty, carts.total_price, carts.users_id, carts.product_id, products.name, products.price, products.category, products.description").Joins("join products on carts.product_id = products.id").Where(where_clause).Find(&cart)

	if query.Error != nil {
		return nil, query.Error
	}
	return cart, nil
}

func CreateCart(Cart *models.Cart) (interface{}, error) {

	if err := config.DB.Create(&Cart).Error; err != nil {
		return nil, err
	}

	return Cart.UsersID, nil
}

// function untuk mendapatkan harga product by id
func GetHargaProduct(id int) (int, error) {
	product := models.Product{}
	err := config.DB.Find(&product, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return product.Price, nil
}

func UpdateCart(id int, Cart *models.Cart) {
	config.DB.Where("id = ?", id).Updates(&Cart)
}

func DeleteCart(id int) (interface{}, error) {
	var cart models.Cart
	check_cart := config.DB.Find(&cart, id).RowsAffected
	err := config.DB.Delete(&cart).Error
	if err != nil || check_cart > 0 {
		return nil, err
	}
	return cart.UsersID, nil
}

// function bantuan untuk mendapatkan is user pada tabel cart
func GetIDUserCart(id int) (uint, uint, error) {
	var cart models.Cart
	err := config.DB.Find(&cart, id)
	if err.Error != nil {
		return 0, 0, err.Error
	}
	return cart.UsersID, cart.ProductID, nil
}
