package controllers

import (
	"ecommerce-api/lib/databases"
	"ecommerce-api/middlewares"
	"ecommerce-api/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateOrderControllers(c echo.Context) error {
	order_req := models.OrderRequest{}
	c.Bind(&order_req)

	var qt, prc int
	for x, v := range order_req.DetailCartId {
		for i := x; i < len(order_req.DetailCartId); i++ {

			log.Println("id detail detail", x)
			id_user_cart, _, _ := databases.GetIDUserCart(v)

			if id_user_cart == 0 {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "Bad Request",
				})
			}
		}
		h, q, _ := databases.GetHargaQtyCart(v)
		qt += q
		prc += h

	}
	log.Println("total qty", qt, " total price:", prc)

	order_req.Order.TotalQty = qt
	order_req.Order.TotalPrice = prc

	logged := middlewares.ExtractTokenId(c)
	order_req.Order.UsersID = uint(logged)

	databases.CreateAddress(&order_req.Address)

	order_req.Order.AddressRequest = order_req.Address.ID
	order_detail, er := databases.CreateOrder(&order_req)
	if er != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	for _, v := range order_req.DetailCartId {
		order := models.ListOrder{}
		order.CartID = uint(v)
		order.OrderID = order_req.Order.ID
		databases.CreateOrderDet(&order)
		databases.DeleteCart(v)
	}

	log.Println("isi cart :", order_req.DetailCartId)
	log.Println("isi city :", order_req.Address.City)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
		"address": order_detail,
	})

}

func GetOrderControllers(c echo.Context) error {

	logged := middlewares.ExtractTokenId(c)
	_, s, err := databases.GetOrder(logged)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
		"address": s,
	})
}
