package helper

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func GetCors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		//AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		//ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}

//http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete

/*
func SetMainMiddlewareForGrpc(e *echo.Echo) {
    e.Use(middleware.RequestID())
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
    }))
}
*/

/*
package helper

import "github.com/go-chi/cors"

func GetCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //ips q tendran permiso
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
*/
