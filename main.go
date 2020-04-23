package main

import (
	"goweb/customer"
	"goweb/database"
	"goweb/orders"
	"goweb/product"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	//dbconection, clientconection := database.LonnectDB()
	dbconection, _ := database.LonnectDB()
	//Product
	var productRepository = product.NewRepository(dbconection)
	var productService = product.NewService(productRepository)
	//customer
	var customerRepository = customer.NewRepository(dbconection)
	var customerService = customer.NewService(customerRepository)
	//customer
	var orderRepository = orders.NewRepository(dbconection)
	var orderService = orders.NewService(orderRepository)
	// Echo instance
	route := echo.New()

	// Middleware
	route.Use(middleware.Logger())
	route.Use(middleware.Recover())

	//define api version
	r := route.Group("/api")
	product.Route(r, productService)
	customer.Route(r, customerService)
	orders.Route(r, orderService)
	//product.Ro(r)

	// Start server
	//database.DisconnectDB(clientconection)
	route.Logger.Fatal(route.Start(":1323"))

}
