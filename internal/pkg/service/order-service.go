package service

import (
	"order-service/internal/pkg/model"
	"order-service/internal/pkg/persistence"
)

type OrderService struct {
	repository persistence.OrderRepository
}

//you can get a pointer using the & operator (e.g. &c)
//you can dereference a pointer to get the value using *
func NewOrderService(repository persistence.OrderRepository) *OrderService {
	return &OrderService{repository: repository}
}

func (s *OrderService) Get(id string) (*model.Order, error) {
	if id == "2" {
		return nil, &MyError{}
	}
	return s.repository.Get(id)
}
