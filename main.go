package main

import (
	"ecommerce-api/config"
	"ecommerce-api/middlewares"
	"ecommerce-api/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))

}
