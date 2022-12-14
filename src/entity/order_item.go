package entity

type OrderItem struct {
	id        string
	productId string
	name      string
	price     float64
	quantity  int64
}

func NewOrderItem(id, productId, name string, price float64, quantity int64) (*OrderItem, error) {
	orderItem := &OrderItem{
		id:        id,
		productId: productId,
		name:      name,
		price:     price,
		quantity:  quantity,
	}

	return orderItem, nil
}

func (o *OrderItem) OrderItemTotal() float64 {
	return o.price * float64(o.quantity)
}
