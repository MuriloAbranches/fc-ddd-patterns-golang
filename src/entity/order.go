package entity

type Order struct {
	id         string
	customerId string
	items      []OrderItem
}

func NewOrder(id, customerId string, items []OrderItem) (*Order, error) {
	order := &Order{
		id:         id,
		customerId: customerId,
		items:      items,
	}

	return order, nil
}
