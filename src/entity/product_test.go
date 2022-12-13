package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductWhenIdIsRequired(t *testing.T) {
	product, err := NewProduct("", "Product 1", 100.00)

	assert.Nil(t, product)
	assert.EqualError(t, ErrorProductIdIsRequired, err.Error())
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("123", "", 100.00)

	assert.Nil(t, product)
	assert.EqualError(t, ErrorProductNameIsRequired, err.Error())
}

func TestProductWhenPriceIsLessThanZero(t *testing.T) {
	product, err := NewProduct("123", "Product 1", -1.0)

	assert.Nil(t, product)
	assert.EqualError(t, ErrorProductPriceIsLessThanZero, err.Error())
}

func TestProductChangeName(t *testing.T) {
	product, _ := NewProduct("123", "Product 1", 100.0)

	err := product.ChangeName("Product 2")

	assert.Nil(t, err)
	assert.Equal(t, "Product 2", product.name)
}

func TestProductChangePrice(t *testing.T) {
	product, _ := NewProduct("123", "Product 1", 100.0)

	err := product.ChangePrice(200.0)

	assert.Nil(t, err)
	assert.Equal(t, 200.0, product.price)
}
