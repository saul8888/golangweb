package authentication

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

//se firman los token con una  llave privada
//create -> openssl genrsa -out private.rsa 1024
//se verifican con una llave publica
//create -> openssl rsa -in private.rsa -pubout > public.rsa.pub

//create variables, type punters for keys

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

type Repository interface {
	GenerateCustomer(params *dateValidate) (*Customer, error)
	GenerateJWT(user Customer) string
	ValidateToken(w *echo.Response, r *http.Request) (*jwt.Token, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(dbconection *mongo.Database) Repository {
	//--------------------init token----------------------//
	//the read archive in format bytes for save  the keys private anad public

	privateBytes, err := ioutil.ReadFile("./private.rsa")
	if err != nil {
		log.Fatal("private key was not read")
	}

	publicBytes, err := ioutil.ReadFile("./public.rsa.pub")
	if err != nil {
		log.Fatal("public key was not read")
	}

	//for load in the form of a key private and public
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("could not do the parse of private")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("could not do the parse of public")
	}

	//----------------------------------------------------//

	return &repository{db: dbconection}
}

func (repo *repository) GenerateCustomer(params *dateValidate) (*Customer, error) {

	collection := repo.db.Collection("user")
	row, err := collection.Find(context.TODO(), bson.M{"email": params.Email, "password": params.Password})
	if err != nil {
		fmt.Println(err)
	}

	customer := &Customer{}
	for row.Next(context.TODO()) {
		row.Decode(&customer)
	}

	return customer, err
}

func (repo *repository) GenerateJWT(user Customer) string {
	//create a struct of my Claim
	claims := Claim{
		Customer: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "token test", //object of token
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	//encode to base64
	result, err := token.SignedString(Keys())
	//result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("could not sign private token")
	}
	return result
}

func (repo *repository) ValidateToken(w *echo.Response, r *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})
	return token, err
}
