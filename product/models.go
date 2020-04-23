package product

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatusAction struct {
	Action string `json:"action"`
	Update int    `json:"update"`
}

type Product struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	ProductName string             `bson:"product_name"`
	Description string             `json:"description"`
	TotalAmount int                `bson:"total_amount"`
	TotalSold   int                `bson:"total_sold"`
	Price       float32            `json:"price"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdateAt    time.Time          `bson:"update_at"`
	//ProductCode string             `json:"product_code"`
}

type ProductList struct {
	Data         []*Product `json:"data"`
	TotalRecords int        `json:"totalRecords"`
}

type NewProduct struct {
	Action string `json:"action"`
	ID     string `json:"newId"`
}
