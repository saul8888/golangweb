package product

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
	GetProductById(id string) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProduct(params *getAddProductRequest) (string, error)
	UpdateProduct(params *updateProductRequest) (int, error)
	DeleteProduct(ProductId string) (int, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(dbconection *mongo.Database) Repository {
	return &repository{db: dbconection}
}

func (repo *repository) GetProductById(id string) (*Product, error) {
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	collection := repo.db.Collection("products")
	row, err := collection.Find(context.TODO(), bson.M{"_id": ObjID})
	if err != nil {
		fmt.Println(err)
	}

	product1 := &Product{}
	product2 := []Product{}

	for row.Next(context.TODO()) {
		row.Decode(&product1)
		product2 = append(product2, *product1)
	}
	return product1, err
}

func (repo *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
	collection := repo.db.Collection("products")
	options := options.Find()
	options.SetSkip(int64(params.Offset))
	options.SetLimit(int64(params.Limit))
	row, err := collection.Find(context.TODO(), bson.M{}, options)
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

func (repo *repository) GetTotalProducts() (int, error) {
	collection := repo.db.Collection("products")
	total, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	return int(total), nil
}

func (repo *repository) InsertProduct(params *getAddProductRequest) (string, error) {
	collection := repo.db.Collection("products")
	// Insert One Document.
	product1 := Product{}
	// An ID for MongoDB.
	product1.ID = primitive.NewObjectID()
	product1.ProductName = params.ProductName
	product1.Description = params.Description
	product1.TotalAmount = params.TotalAmount
	product1.TotalSold = params.TotalSold
	product1.Price = params.Price
	product1.CreatedAt = time.Now()
	product1.UpdateAt = time.Now()

	hola, err := collection.InsertOne(context.TODO(), product1)
	fmt.Println(hola)
	if err != nil {
		panic(err)
	}
	return product1.ID.String(), nil

}

func (repo *repository) UpdateProduct(param *updateProductRequest) (int, error) {
	collection := repo.db.Collection("products")

	objID, err := primitive.ObjectIDFromHex(param.ID)
	if err != nil {
		panic(err)
	}
	resultUpdate, err := collection.UpdateOne(context.TODO(),
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"product_name": param.ProductName,
				"description":  param.Description,
				"total_amount": param.TotalAmount,
				"total_sold":   param.TotalSold,
				"price":        param.Price,
				"update_at":    time.Now(),
			},
		},
	)

	return int(resultUpdate.ModifiedCount), nil // output: 1

}

func (repo *repository) DeleteProduct(ProductId string) (int, error) {
	collection := repo.db.Collection("products")

	objID, err := primitive.ObjectIDFromHex(ProductId)
	if err != nil {
		panic(err)
	}
	resultDelete, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		fmt.Println(err)
	}

	return int(resultDelete.DeletedCount), nil
}
