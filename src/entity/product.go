package entity

import "errors"

var (
	ErrorProductIdIsRequired        = errors.New("id is required")
	ErrorProductNameIsRequired      = errors.New("name is required")
	ErrorProductPriceIsLessThanZero = errors.New("price must be greater than zero")
)

type Product struct {
	id    string
	name  string
	price float64
}

func NewProduct(id, name string, price float64) (*Product, error) {
	product := &Product{
		id:    id,
		name:  name,
		price: price,
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) ChangeName(name string) error {
	p.name = name
	return p.Validate()
}

func (p *Product) ChangePrice(price float64) error {
	p.price = price
	return p.Validate()
}

func (p *Product) Validate() error {
	if p.id == "" {
		return ErrorProductIdIsRequired
	}

	if p.name == "" {
		return ErrorProductNameIsRequired
	}

	if p.price < 0 {
		return ErrorProductPriceIsLessThanZero
	}

	return nil
}

func (p *Product) GetPrice() float64 {
	return p.price
}
