package entity

import "errors"

var (
	ErrorCustomerNameIsRequired     = errors.New("name is required")
	ErrorCustomerIdIsRequired       = errors.New("id is required")
	ErrorCustomerAddressIsMandatory = errors.New("address is mandatory to activate a customer")
)

type Customer struct {
	id           string
	name         string
	address      Address
	active       bool
	rewardPoints int
}

func NewCustomer(id, name string) (*Customer, error) {
	customer := &Customer{
		id:           id,
		name:         name,
		address:      Address{},
		active:       false,
		rewardPoints: 0,
	}

	err := customer.Validate()
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) Validate() error {
	if c.name == "" {
		return ErrorCustomerNameIsRequired
	}

	if c.id == "" {
		return ErrorCustomerIdIsRequired
	}

	return nil
}

func (c *Customer) ChangeName(name string) error {
	c.name = name
	return c.Validate()
}

func (c *Customer) IsActive() bool {
	return c.active
}

func (c *Customer) Activate() error {
	if c.address == (Address{}) {
		return ErrorCustomerAddressIsMandatory
	}

	c.active = true

	return nil
}

func (c *Customer) Deactivate() {
	c.active = false
}

func (c *Customer) AddRewardPoints(points int) {
	c.rewardPoints += points
}

func (c *Customer) GetRewardPoints() int {
	return c.rewardPoints
}

func (c *Customer) GetId() string {
	return c.id
}

func (c *Customer) SetAddress(address Address) {
	c.address = address
}
