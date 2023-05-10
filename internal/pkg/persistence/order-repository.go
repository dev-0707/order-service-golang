package persistence

import (
	"order-service/internal/pkg/model"
)

type OrderRepository struct{}

/*
s := "hello"      // type string
t := "bye"        // type string
u := 44           // type int
v := [2]int{1, 2} // type array
q := &s           // type pointer
var p *string = &s // type *string
q := &s // shorthand, same dea
*/
var orderRepository *OrderRepository

func GetOrderRepository() *OrderRepository {
	if orderRepository == nil {
		orderRepository = &OrderRepository{}
	}
	return orderRepository
}

func (r *OrderRepository) Get(id string) (*model.Order, error) {
	var order model.Order
	where := model.Order{}
	where.Id = id
	notFound, err := First(&where, &order, []string{})

	if notFound {
		return nil, err
	}

	return &order, err
}
