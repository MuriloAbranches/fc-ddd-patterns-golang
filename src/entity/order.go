package entity

import "errors"

var (
	ErrorOrderIdIsRequired         = errors.New("id is required")
	ErrorOrderCustomerIdIsRequired = errors.New("customerId is required")
	ErrorOrderItemsAreRequired     = errors.New("items are required")
)

type Order struct {
	id         string
	customerId string
	items      []OrderItem
	total      float64
}

func NewOrder(id, customerId string, items []OrderItem) (*Order, error) {
	order := &Order{
		id:         id,
		customerId: customerId,
		items:      items,
	}

	order.total = order.Total()
	err := order.Validate()
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) Validate() error {
	if o.id == "" {
		return ErrorOrderIdIsRequired
	}

	if o.customerId == "" {
		return ErrorOrderCustomerIdIsRequired
	}

	if len(o.items) == 0 {
		return ErrorOrderItemsAreRequired
	}

	return nil
}

func (o *Order) Total() float64 {
	var total = 0.0

	for _, orderItem := range o.items {
		total = total + orderItem.price
	}

	return total
}
