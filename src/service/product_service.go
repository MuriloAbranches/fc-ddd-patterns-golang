package service

import (
	"github.com/MuriloAbranches/fc-ddd-patterns-golang/src/entity"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (p *ProductService) IncreasePrice(products []entity.Product, percentage float64) []entity.Product {
	for i, product := range products {
		products[i].ChangePrice((product.GetPrice()*percentage)/100 + product.GetPrice())
	}

	return products
}
