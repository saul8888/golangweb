package customer

import (
	"net/http"

	"github.com/labstack/echo"
)

type Service interface {
	GetCustomerById(context echo.Context) error
	GetCustomers(context echo.Context) error
	InsertCustomer(context echo.Context) error
	UpdateCustomer(context echo.Context) error
	DeleteCustomer(context echo.Context) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetCustomerById(context echo.Context) error {
	customerID := context.QueryParam("ID")
	customer, err := s.repo.GetCustomerById(customerID)
	if err != nil {
		panic(err)
	}
	return context.JSON(http.StatusOK, customer)
}

func (s *service) GetCustomers(context echo.Context) (err error) {
	params := new(getCustomersRequest)
	if err = context.Bind(params); err != nil {
		return
	}
	customers, err := s.repo.GetCustomers(params)
	if err != nil {
		panic(err)
	}
	totalCustomers, err := s.repo.GetTotalCustomers()
	if err != nil {
		panic(err)
	}
	cant := CustomerList{Data: customers, TotalRecords: totalCustomers}
	return context.JSON(http.StatusOK, cant)

}

func (s *service) InsertCustomer(context echo.Context) (err error) {
	customer := new(getAddCustomerRequest)
	if err = context.Bind(customer); err != nil {
		return
	}
	newCustomer, err := s.repo.InsertCustomer(customer)
	if err != nil {
		panic(err)
	}
	result := NewCustomer{Action: "OK", ID: newCustomer}
	return context.JSON(http.StatusOK, result)
}

func (s *service) UpdateCustomer(context echo.Context) (err error) {
	customer := new(updateCustomerRequest)
	if err = context.Bind(customer); err != nil {
		return
	}
	updateCustomer, err := s.repo.UpdateCustomer(customer)
	if err != nil {
		panic(err)
	}
	result := StatusAction{Action: "OK", Update: updateCustomer}
	return context.JSON(http.StatusOK, result)
}

func (s *service) DeleteCustomer(context echo.Context) error {
	customerID := context.QueryParam("ID")
	del, err := s.repo.DeleteCustomer(customerID)
	if err != nil {
		panic(err)
	}

	result := StatusAction{Action: "OK", Update: del}
	return context.JSON(http.StatusOK, result)
}
