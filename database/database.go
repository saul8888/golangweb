package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoCustomerRepo stores database connection information
type mongodata struct {
	db     *mongo.Database
	client *mongo.Client
}

//------------------Connect method------------------//
func LonnectDB() (*mongo.Database, *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() //will run when  we're finished main

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://saul:1234@cluster0-ooeaq.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	client.Connect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("eatos")

	return db, client
}

//------------------Disconnect method------------------//
func DisconnectDB(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Disconnect(ctx)

	return err
}

/*
func (mongodb *MongoData) GetAll() ([]*Product, error) {
	customersCollection := mongodb.db.Collection("products")
	options := options.Find()
	options.SetSkip(0)
	//options.SetSkip(int64(params.Offset))
	options.SetLimit(1)
	//options.SetLimit(int64(params.Limit))
	row, err := customersCollection.Find(mongodb.context, bson.M{}, options)

	if err != nil {
		fmt.Println(err)
	}

	var product2 []*Product

	for row.Next(context.TODO()) {
		product := &Product{}
		row.Decode(&product)
		product2 = append(product2, product)
	}

	return product2, err
}
*/
