package entity

type OrderItem struct {
	id    string
	name  string
	price float64
}

func NewOrderItem(id, name string, price float64) (*OrderItem, error) {
	orderItem := &OrderItem{
		id:    id,
		name:  name,
		price: price,
	}

	return orderItem, nil
}
