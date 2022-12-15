package service

import (
	"testing"

	"github.com/MuriloAbranches/fc-ddd-patterns-golang/src/entity"
	"github.com/stretchr/testify/assert"
)

func TestProductServiceChangePrice(t *testing.T) {
	product1, _ := entity.NewProduct("product1", "Product 1", 10.0)
	product2, _ := entity.NewProduct("product2", "Product 2", 20.0)

	products := []entity.Product{*product1, *product2}

	productService := NewProductService()
	products = productService.IncreasePrice(products, 100.0)

	assert.Equal(t, 20.0, products[0].GetPrice())
	assert.Equal(t, 40.0, products[1].GetPrice())
}
