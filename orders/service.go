package orders

import (
	"net/http"

	"github.com/labstack/echo"
)

type Service interface {
	GetOrderById(context echo.Context) error
	GetOrders(context echo.Context) error
	InsertOrder(context echo.Context) error
	UpdateOrder(context echo.Context) error
	DeleteOrder(context echo.Context) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetOrderById(context echo.Context) error {
	orderID := context.QueryParam("ID")
	order, err := s.repo.GetOrderById(orderID)
	if err != nil {
		panic(err)
	}
	return context.JSON(http.StatusOK, order)
}

func (s *service) GetOrders(context echo.Context) (err error) {
	params := new(getOrdersRequest)
	if err = context.Bind(params); err != nil {
		return
	}
	orders, err := s.repo.GetOrders(params)
	if err != nil {
		panic(err)
	}
	totalOrders, err := s.repo.GetTotalOrders()
	if err != nil {
		panic(err)
	}
	cant := OrderList{Data: orders, TotalRecords: totalOrders}
	return context.JSON(http.StatusOK, cant)

}

func (s *service) InsertOrder(context echo.Context) (err error) {
	order := new(getAddOrderRequest)
	if err = context.Bind(order); err != nil {
		return
	}
	newOrder, err := s.repo.InsertOrder(order)
	if err != nil {
		panic(err)
	}
	result := NewOrder{Action: "OK", ID: newOrder}
	return context.JSON(http.StatusOK, result)
}

func (s *service) UpdateOrder(context echo.Context) (err error) {
	order := new(updateOrderRequest)
	if err = context.Bind(order); err != nil {
		return
	}
	updateOrder, err := s.repo.UpdateOrder(order)
	if err != nil {
		panic(err)
	}
	result := StatusAction{Action: "OK", Update: updateOrder}
	return context.JSON(http.StatusOK, result)
}

func (s *service) DeleteOrder(context echo.Context) error {
	orderID := context.QueryParam("ID")
	del, err := s.repo.DeleteOrder(orderID)
	if err != nil {
		panic(err)
	}

	result := StatusAction{Action: "OK", Update: del}
	return context.JSON(http.StatusOK, result)
}
