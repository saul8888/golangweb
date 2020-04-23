package customer

import (
	"github.com/labstack/echo"
)

// Register a new user
func Route(r *echo.Group, s Service) {
	customer := r.Group("/customers")
	// Routes
	customer.GET("/get1", s.GetCustomerById)
	customer.GET("/get2", s.GetCustomers)
	customer.POST("/post", s.InsertCustomer)
	customer.PUT("/put", s.UpdateCustomer)
	customer.DELETE("/delete", s.DeleteCustomer)

}
