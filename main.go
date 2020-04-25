package main

import (
	"goweb/authentication"
	"goweb/customer"
	"goweb/database"
	"goweb/helper"
	"goweb/product"
	"net/http"

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

	// Configure middleware with the custom claims type
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &authentication.Claim{},
		SigningKey: authentication.Keys(),
	}))
	//r.POST("/hola", hola, authentication.ConfigToken)
	product.Route(r, productService)
	customer.Route(r, customerService)

	// Start server
	//database.DisconnectDB(clientconection)
	route.Logger.Fatal(route.Start(":8000"))

}

func hola(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome hola")
}
