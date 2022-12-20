package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerWhenIdIsRequired(t *testing.T) {
	customer, err := NewCustomer("", "Customer 1")

	assert.Nil(t, customer)
	assert.EqualError(t, ErrorCustomerIdIsRequired, err.Error())
}

func TestCustomerWhenNameIsRequired(t *testing.T) {
	customer, err := NewCustomer("123", "")

	assert.Nil(t, customer)
	assert.EqualError(t, ErrorCustomerNameIsRequired, err.Error())
}

func TestCustomerChangeName(t *testing.T) {
	customer, err := NewCustomer("123", "John")

	customer.ChangeName("Jane")

	assert.Nil(t, err)
	assert.Equal(t, "Jane", customer.name)
}

func TestCustomerActivate(t *testing.T) {
	customer, _ := NewCustomer("123", "Customer 1")
	address, _ := NewAddress("Street 1", 123, "123456-789", "São Paulo")
	customer.SetAddress(*address)

	err := customer.Activate()

	assert.Nil(t, err)
	assert.True(t, customer.IsActive())
}

func TestCustomerDeactivate(t *testing.T) {
	customer, _ := NewCustomer("123", "Customer 1")

	customer.Deactivate()

	assert.False(t, customer.IsActive())
}

func TestCustomerWhenAddressIsMandatory(t *testing.T) {
	customer, _ := NewCustomer("123", "Customer 1")

	err := customer.Activate()

	assert.EqualError(t, ErrorCustomerAddressIsMandatory, err.Error())
	assert.False(t, customer.IsActive())
}

func TestCustomerAddRewardPoints(t *testing.T) {
	customer, err := NewCustomer("1", "Customer 1")
	assert.Nil(t, err)
	assert.Equal(t, 0, customer.GetRewardPoints())

	customer.AddRewardPoints(10)

	assert.Equal(t, 10, customer.GetRewardPoints())

	customer.AddRewardPoints(10)

	assert.Equal(t, 20, customer.GetRewardPoints())
}
