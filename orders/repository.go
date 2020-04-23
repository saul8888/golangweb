package orders

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	GetOrderById(id string) (*Order, error)
	GetOrders(params *getOrdersRequest) ([]*Order, error)
	GetTotalOrders() (int, error)
	InsertOrder(params *getAddOrderRequest) (string, error)
	UpdateOrder(params *updateOrderRequest) (int, error)
	DeleteOrder(OrderId string) (int, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(dbconection *mongo.Database) Repository {
	return &repository{db: dbconection}
}

func (repo *repository) GetOrderById(id string) (*Order, error) {
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	collection := repo.db.Collection("user")
	row, err := collection.Find(context.TODO(), bson.M{"_id": ObjID})
	if err != nil {
		fmt.Println(err)
	}

	order1 := &Order{}
	order2 := []Order{}

	for row.Next(context.TODO()) {
		row.Decode(&order1)
		order2 = append(order2, *order1)
	}
	return order1, err
}

func (repo *repository) GetOrders(params *getOrdersRequest) ([]*Order, error) {
	collection := repo.db.Collection("user")
	options := options.Find()
	options.SetSkip(int64(params.Offset))
	options.SetLimit(int64(params.Limit))
	row, err := collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		fmt.Println(err)
	}

	var order2 []*Order

	for row.Next(context.TODO()) {
		order := &Order{}
		row.Decode(&order)
		order2 = append(order2, order)
	}

	return order2, err

}

func (repo *repository) GetTotalOrders() (int, error) {
	collection := repo.db.Collection("order")
	total, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	return int(total), nil
}

func (repo *repository) InsertOrder(params *getAddOrderRequest) (string, error) {
	collection := repo.db.Collection("user")
	// Insert One Document.
	order1 := Order{}
	// An ID for MongoDB.
	order1.ID = primitive.NewObjectID()
	order1.Name = params.Name
	order1.Email = params.Email
	order1.PhoneNumber = params.PhoneNumber
	order1.Password = params.Password
	order1.CreatedAt = time.Now()
	order1.UpdateAt = time.Now()

	newOrder, err := collection.InsertOne(context.TODO(), order1)
	//fmt.Println(newOrder.InsertedID)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v", newOrder.InsertedID), nil //return order1.ID.String()

}

func (repo *repository) UpdateOrder(param *updateOrderRequest) (int, error) {
	collection := repo.db.Collection("user")

	objID, err := primitive.ObjectIDFromHex(param.ID)
	if err != nil {
		panic(err)
	}
	resultUpdate, err := collection.UpdateOne(context.TODO(),
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"customer_name": param.Name,
				"description":   param.Email,
				"total_amount":  param.PhoneNumber,
				"password":      param.Password,
				"update_at":     time.Now(),
			},
		},
	)

	return int(resultUpdate.ModifiedCount), nil // output: 1

}

func (repo *repository) DeleteOrder(OrderId string) (int, error) {
	collection := repo.db.Collection("user")

	objID, err := primitive.ObjectIDFromHex(OrderId)
	if err != nil {
		panic(err)
	}
	resultDelete, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		fmt.Println(err)
	}

	return int(resultDelete.DeletedCount), nil
}
