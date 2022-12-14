package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderWhenIdIsRequired(t *testing.T) {
	order, err := NewOrder("", "123", make([]OrderItem, 0))

	assert.Nil(t, order)
	assert.EqualError(t, ErrorOrderIdIsRequired, err.Error())
}

func TestOrderWhenCustomerIdIsRequired(t *testing.T) {
	order, err := NewOrder("123", "", make([]OrderItem, 0))

	assert.Nil(t, order)
	assert.EqualError(t, ErrorOrderCustomerIdIsRequired, err.Error())
}

func TestOrderWhenOrderItemsAreRequired(t *testing.T) {
	order, err := NewOrder("123", "123", make([]OrderItem, 0))

	assert.Nil(t, order)
	assert.EqualError(t, ErrorOrderItemsAreRequired, err.Error())
}

func TestOrderCalculateTotal(t *testing.T) {
	item1, _ := NewOrderItem("i1", "p1", "Item 1", 100.00, 2)
	order1, err := NewOrder("o1", "c1", []OrderItem{*item1})

	assert.Nil(t, err)
	assert.Equal(t, 200.00, order1.total)

	item2, _ := NewOrderItem("i2", "p2", "Item 2", 200.00, 2)
	order2, err := NewOrder("o2", "c2", []OrderItem{*item1, *item2})

	assert.Nil(t, err)
	assert.Equal(t, 600.00, order2.total)
}

func TestOrderWhenItemQuantityIsLessOrEqualZero(t *testing.T) {
	item1, _ := NewOrderItem("i1", "p1", "Item 1", 100.00, 0)
	order1, err := NewOrder("o1", "c1", []OrderItem{*item1})

	assert.Nil(t, order1)
	assert.EqualError(t, ErrorItemQuantityIsLessOrEqualZero, err.Error())
}
