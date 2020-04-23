package orders

import (
	"github.com/labstack/echo"
)

// Register a new user
func Route(r *echo.Group, s Service) {
	order := r.Group("/orders")
	// Routes
	order.GET("/get1", s.GetOrderById)
	order.GET("/get2", s.GetOrders)
	order.POST("/post", s.InsertOrder)
	order.PUT("/put", s.UpdateOrder)
	order.DELETE("/delete", s.DeleteOrder)

}
