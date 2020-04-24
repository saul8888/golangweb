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
	} else {
		token := s.repo.GenerateJWT(*customer)
		result := Responsetoken{Token: token}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			fmt.Println("error generating the json")
			//return
		}
		context.Response().WriteHeader(http.StatusOK)
		context.Response().Header().Set("Content-Type", "application/json")
		//responseJSON(w, newUser)
		context.Response().Write(jsonResult)
	}

	//fmt.Println(*customer)
	return context.JSON(http.StatusOK, customer)
}

func (s *service) ValidateToken(context echo.Context) error {
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
}
