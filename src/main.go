package main

import (
	"fmt"

	"github.com/MuriloAbranches/fc-ddd-patterns-golang/src/entity"
)

func main() {

	customer, err := entity.NewCustomer("123", "Murilo Abranches")
	if err != nil {
		panic(err)
	}

	address, err := entity.NewAddress("Rua Um", 1, "123456-789", "São Paulo")
	if err != nil {
		panic(err)
	}

	fmt.Println(address.ToString())

	customer.SetAddress(*address)
	customer.Activate()

	fmt.Println(customer)

	item1, err := entity.NewOrderItem("1", "p1", "Item 1", 10, 2)
	if err != nil {
		panic(err)
	}

	item2, err := entity.NewOrderItem("2", "p2", "Item 2", 15, 1)
	if err != nil {
		panic(err)
	}

	items := []entity.OrderItem{*item1, *item2}

	order, err := entity.NewOrder("1", "123", items)
	if err != nil {
		panic(err)
	}

	fmt.Println(order)
}
