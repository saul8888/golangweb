package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Service interface {
	GenerateCustomer(context echo.Context) error
	ValidateToken(context echo.Context) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GenerateCustomer(context echo.Context) (err error) {
	date := new(dateValidate)
	if err = context.Bind(date); err != nil {
		return
	}
	customer, err := s.repo.GenerateCustomer(date)
	if err != nil {
		panic(err)
	}
	if customer.Name != "" && customer.Email != "" {
		token := s.repo.GenerateJWT(*customer)
		result := Responsetoken{Token: token}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			fmt.Println("error generating the json")
			//return
		}
		context.Response().WriteHeader(http.StatusOK)
		context.Response().Header().Set("Content-Type", "application/json")
		context.Response().Write(jsonResult)
	} else {
		context.Response().WriteHeader(http.StatusForbidden)
		result := Responsetoken{Token: "usser or password invalid"}
		jsonResult, _ := json.Marshal(result)
		fmt.Println(context.Response(), "usser or password invalid")
		context.Response().Write(jsonResult)
	}

	//return context.JSON(http.StatusOK, jsonResult)
	return nil
}

func (s *service) ValidateToken(context echo.Context) error {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(*Claim)
	name := claims.Name
	return context.String(http.StatusOK, "Welcome "+name+"!")
	/*
		token, err := s.repo.ValidateToken(context.Response(), context.Request())
		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError:
				vErr := err.(*jwt.ValidationError)
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					fmt.Fprintln(context.Response(), "your token expired")
					//return
				case jwt.ValidationErrorSignatureInvalid:
					fmt.Fprintln(context.Response(), "the signature does not match")
					//return
				default:
					fmt.Fprintln(context.Response(), "the signature does not match")
					//return
				}
			default:
				fmt.Fprintln(context.Response(), "your token is not valid")
				//return
			}
		}
		if token.Valid {
			context.Response().WriteHeader(http.StatusAccepted)
			fmt.Fprintln(context.Response(), "welcome to the system")
		} else {
			context.Response().WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(context.Response(), "your token is not valid")
		}
		return nil
	*/
}
