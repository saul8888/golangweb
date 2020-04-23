package product

import (
	"github.com/labstack/echo"
)

// Register a new user
func Route(r *echo.Group, s Service) {
	product := r.Group("/products")
	// Routes
	product.GET("/get1", s.GetProductById)
	product.GET("/get2", s.GetProducts)
	product.POST("/post", s.InsertProduct)
	product.PUT("/put", s.UpdateProduct)
	product.DELETE("/delete", s.DeleteProduct)

}
