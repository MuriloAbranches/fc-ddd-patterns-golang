package service

import (
	"errors"

	"github.com/MuriloAbranches/fc-ddd-patterns-golang/src/entity"
	"github.com/google/uuid"
)

var (
	ErrorOrderMustHaveOneItem = errors.New("order must have at least one item")
)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (os *OrderService) PlaceOrder(customer *entity.Customer, items []entity.OrderItem) (*entity.Order, error) {

	if len(items) == 0 {
		return nil, ErrorOrderMustHaveOneItem
	}

	order, err := entity.NewOrder(uuid.New().String(), customer.GetId(), items)
	if err != nil {
		return nil, err
	}

	customer.AddRewardPoints(int(order.Total()) / 2)

	return order, nil
}

func (os *OrderService) Total(orders []entity.Order) float64 {
	total := 0.0

	for _, order := range orders {
		total = total + order.Total()
	}

	return total
}
