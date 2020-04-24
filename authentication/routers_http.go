package authentication

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// Register a new user
func Route(r *echo.Group, s Service) {
	// Routess.
	r.POST("/authentication", s.GenerateCustomer)
	r.POST("/validate", s.ValidateToken)
	r.POST("/aut", authentication)

}

func authentication(context echo.Context) error {
	fmt.Println("###############################3")
	aa := context.Response()
	fmt.Printf("%T", aa)
	fmt.Println()
	bb := context.Request()
	fmt.Printf("%T", bb)
	fmt.Println()
	fmt.Println("###############################3")
	return context.String(http.StatusOK, "Hello, World!")
}
