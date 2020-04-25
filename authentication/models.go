package authentication

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/labstack/echo/middleware"
)

type dateValidate struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

type Customer struct {
	Name      string    `bson:"customer_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `bson:"created_at"`
}

func Keys() []byte {
	privateBytes, err := ioutil.ReadFile("./private.rsa")
	if err != nil {
		log.Fatal("private key was not read")
	}
	return privateBytes
}

var ConfigToken = middleware.JWTWithConfig(middleware.JWTConfig{
	Claims: &Claim{},
	//SigningMethod: "RS256",
	//SigningKey:    publicKey,
	SigningKey: Keys(),
	//SigningMethod: "HS512",//"RS256"
	//TokenLookup:   "header:Authorization"}))

})
