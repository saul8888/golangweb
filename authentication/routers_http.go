package authentication

import (
	"github.com/labstack/echo"
)

// Register a new user
func Route(r *echo.Group, s Service) {
	// Routess.
	r.POST("/authentication", s.GenerateCustomer)
	r.POST("/validate", s.ValidateToken)

}
