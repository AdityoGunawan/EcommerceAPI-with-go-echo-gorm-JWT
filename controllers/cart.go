package controllers

import (
	"ecommerce-api/lib/databases"
	"ecommerce-api/middlewares"
	"ecommerce-api/models"
	"ecommerce-api/response"
	"net/http"
	"strconv"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateCartControllers(c echo.Context) error {

	Cart := models.Cart{}
	c.Bind(&Cart)
	v := validator.New()
	e := v.Var(Cart.Qty, "required,gt=0")
	if e == nil {
		logged := middlewares.ExtractTokenId(c)

		id_user_cart, _ := databases.GetIDUserProduct(int(Cart.ProductID))
		product_price, _ := databases.GetHargaProduct(int(Cart.ProductID))

		Cart.UsersID = uint(logged)
		Cart.TotalPrice = Cart.Qty * product_price

		if uint(logged) == id_user_cart {
			return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
		}
		_, e = databases.CreateCart(&Cart)
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// Get all cart
func GetAllCartControllers(c echo.Context) error {
	logged := middlewares.ExtractTokenId(c)
	cart, err := databases.GetAllCart(logged)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(cart))
}

// Delete cart by id
func DeleteCartControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	id_user_cart, _, _ := databases.GetIDUserCart(conv_id)
	logged := middlewares.ExtractTokenId(c)
	if uint(logged) != id_user_cart {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	databases.DeleteCart(conv_id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func UpdateCartControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}

	cart := models.Cart{}
	c.Bind(&cart)
	v := validator.New()
	e := v.Var(cart.Qty, "required,gt=0")
	if e == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	// mengecek user id nya sama dan ada pada tabel
	id_user_cart, id_product, _ := databases.GetIDUserCart(id)
	logged := middlewares.ExtractTokenId(c)
	if id_user_cart == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	} else if uint(logged) != id_user_cart {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}

	// mengupdate total harga
	product_price, _ := databases.GetHargaProduct(int(id_product))
	cart.TotalPrice = cart.Qty * product_price

	// untuk mengupdate
	databases.UpdateCart(id, &cart)

	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
