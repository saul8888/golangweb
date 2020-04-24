package authentication

import (
	"time"
)

type dateValidate struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

type Customer struct {
	Name      string    `bson:"customer_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `bson:"created_at"`
	//ProductCode string             `json:"product_code"`
}
