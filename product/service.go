package product

import (
	"net/http"

	"github.com/labstack/echo"
)

type Service interface {
	GetProductById(context echo.Context) error
	GetProducts(context echo.Context) error
	InsertProduct(context echo.Context) error
	UpdateProduct(context echo.Context) error
	DeleteProduct(context echo.Context) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetProductById(context echo.Context) error {
	customerID := context.QueryParam("ID")
	product, err := s.repo.GetProductById(customerID)
	if err != nil {
		panic(err)
	}
	return context.JSON(http.StatusOK, product)
}

func (s *service) GetProducts(context echo.Context) (err error) {
	params := new(getProductsRequest)
	if err = context.Bind(params); err != nil {
		return
	}
	products, err := s.repo.GetProducts(params)
	if err != nil {
		panic(err)
	}
	totalProducts, err := s.repo.GetTotalProducts()
	if err != nil {
		panic(err)
	}
	cant := ProductList{Data: products, TotalRecords: totalProducts}
	return context.JSON(http.StatusOK, cant)

}

func (s *service) InsertProduct(context echo.Context) (err error) {
	product := new(getAddProductRequest)
	if err = context.Bind(product); err != nil {
		return
	}
	newProduct, err := s.repo.InsertProduct(product)
	if err != nil {
		panic(err)
	}
	result := NewProduct{Action: "OK", ID: newProduct}
	return context.JSON(http.StatusOK, result)
}

func (s *service) UpdateProduct(context echo.Context) (err error) {
	product := new(updateProductRequest)
	if err = context.Bind(product); err != nil {
		return
	}
	updateProduct, err := s.repo.UpdateProduct(product)
	if err != nil {
		panic(err)
	}
	result := StatusAction{Action: "OK", Update: updateProduct}
	return context.JSON(http.StatusOK, result)
}

func (s *service) DeleteProduct(context echo.Context) error {
	customerID := context.QueryParam("ID")
	del, err := s.repo.DeleteProduct(customerID)
	if err != nil {
		panic(err)
	}

	result := StatusAction{Action: "OK", Update: del}
	return context.JSON(http.StatusOK, result)
}
