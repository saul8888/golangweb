package orders

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatusAction struct {
	Action string `json:"action"`
	Update int    `json:"update"`
}

type Order struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name        string             `bson:"customer_name"`
	Email       string             `json:"email"`
	PhoneNumber string             `bson:"phone_number"`
	Password    string             `bson:"password"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdateAt    time.Time          `bson:"update_at"`
	//ProductCode string             `json:"product_code"`
}

type OrderList struct {
	Data         []*Order `json:"data"`
	TotalRecords int      `json:"totalRecords"`
}

type NewOrder struct {
	Action string `json:"action"`
	ID     string `json:"newId"`
}
