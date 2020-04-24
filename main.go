package main

import (
	"goweb/authentication"
	"goweb/customer"
	"goweb/database"
	"goweb/helper"
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
	//orders
	//var orderRepository = orders.NewRepository(dbconection)
	//var orderService = orders.NewService(orderRepository)
	//authentication
	var autheRepository = authentication.NewRepository(dbconection)
	var autheService = authentication.NewService(autheRepository)
	// Echo instance
	route := echo.New()

	// Middleware
	route.Use(middleware.Logger())
	route.Use(middleware.Recover())
	route.Use(helper.GetCors())

	//define api version
	r := route.Group("/api")
	authentication.Route(r, autheService)
	r.Use(middleware.BasicAuth(he))
	product.Route(r, productService)
	customer.Route(r, customerService)
	//product.Ro(r)

	// Start server
	//database.DisconnectDB(clientconection)
	route.Logger.Fatal(route.Start(":1323"))

}

func he(username, password string, c echo.Context) (bool, error) {
	if username == "sa" && password == "sa" {
		return true, nil
	}
	return false, nil
}
