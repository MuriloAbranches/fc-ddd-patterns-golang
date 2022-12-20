package service

import (
	"testing"

	"github.com/MuriloAbranches/fc-ddd-patterns-golang/src/entity"
	"github.com/stretchr/testify/assert"
)

func TestOrderServicePlaceOrder(t *testing.T) {
	customer, _ := entity.NewCustomer("c1", "Customer 1")
	item1, _ := entity.NewOrderItem("i1", "p1", "Item 1", 10.0, 1)

	orderService := NewOrderService()
	order, err := orderService.PlaceOrder(customer, []entity.OrderItem{*item1})

	assert.Nil(t, err)
	assert.Equal(t, 5, customer.GetRewardPoints())
	assert.Equal(t, 10.0, order.Total())
}

func TestOrderServiceGetTotalOfAllOrders(t *testing.T) {
	item1, _ := entity.NewOrderItem("i1", "p1", "Item 1", 100.0, 1)
	item2, _ := entity.NewOrderItem("i2", "p2", "Item 2", 200.0, 2)

	order1, _ := entity.NewOrder("o1", "c1", []entity.OrderItem{*item1})
	order2, _ := entity.NewOrder("o2", "c2", []entity.OrderItem{*item2})

	orderService := NewOrderService()
	total := orderService.Total([]entity.Order{*order1, *order2})

	assert.Equal(t, 500.0, total)
}
