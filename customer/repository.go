package customer

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
	GetCustomerById(id string) (*Customer, error)
	GetCustomers(params *getCustomersRequest) ([]*Customer, error)
	GetTotalCustomers() (int, error)
	InsertCustomer(params *getAddCustomerRequest) (string, error)
	UpdateCustomer(params *updateCustomerRequest) (int, error)
	DeleteCustomer(CustomerId string) (int, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(dbconection *mongo.Database) Repository {
	return &repository{db: dbconection}
}

func (repo *repository) GetCustomerById(id string) (*Customer, error) {
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	collection := repo.db.Collection("user")
	row, err := collection.Find(context.TODO(), bson.M{"_id": ObjID})
	if err != nil {
		fmt.Println(err)
	}

	customer1 := &Customer{}
	customer2 := []Customer{}

	for row.Next(context.TODO()) {
		row.Decode(&customer1)
		customer2 = append(customer2, *customer1)
	}
	return customer1, err
}

func (repo *repository) GetCustomers(params *getCustomersRequest) ([]*Customer, error) {
	collection := repo.db.Collection("user")
	options := options.Find()
	options.SetSkip(int64(params.Offset))
	options.SetLimit(int64(params.Limit))
	row, err := collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		fmt.Println(err)
	}

	var customer2 []*Customer

	for row.Next(context.TODO()) {
		customer := &Customer{}
		row.Decode(&customer)
		customer2 = append(customer2, customer)
	}

	return customer2, err

}

func (repo *repository) GetTotalCustomers() (int, error) {
	collection := repo.db.Collection("user")
	total, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	return int(total), nil
}

func (repo *repository) InsertCustomer(params *getAddCustomerRequest) (string, error) {
	collection := repo.db.Collection("user")
	// Insert One Document.
	customer1 := Customer{}
	// An ID for MongoDB.
	customer1.ID = primitive.NewObjectID()
	customer1.Name = params.Name
	customer1.Email = params.Email
	customer1.PhoneNumber = params.PhoneNumber
	customer1.Password = params.Password
	customer1.CreatedAt = time.Now()
	customer1.UpdateAt = time.Now()

	newCustomer, err := collection.InsertOne(context.TODO(), customer1)
	//fmt.Println(newCustomer.InsertedID)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v", newCustomer.InsertedID), nil //return customer1.ID.String()

}

func (repo *repository) UpdateCustomer(param *updateCustomerRequest) (int, error) {
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

func (repo *repository) DeleteCustomer(CustomerId string) (int, error) {
	collection := repo.db.Collection("user")

	objID, err := primitive.ObjectIDFromHex(CustomerId)
	if err != nil {
		panic(err)
	}
	resultDelete, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		fmt.Println(err)
	}

	return int(resultDelete.DeletedCount), nil
}
