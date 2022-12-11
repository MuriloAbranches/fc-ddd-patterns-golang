package entity

import "errors"

var (
	ErrorNameIsRequired     = errors.New("name is required")
	ErrorIdIsRequired       = errors.New("id is required")
	ErrorAddressIsMandatory = errors.New("address is mandatory to activate a customer")
)

type Customer struct {
	id      string
	name    string
	address Address
	active  bool
}

func NewCustomer(id, name string) (*Customer, error) {
	customer := &Customer{
		id:      id,
		name:    name,
		address: Address{},
		active:  false,
	}

	err := customer.Validate()
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) Validate() error {
	if c.name == "" {
		return ErrorNameIsRequired
	}

	if c.id == "" {
		return ErrorIdIsRequired
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
		return ErrorAddressIsMandatory
	}

	c.active = true

	return nil
}

func (c *Customer) Deactivate() {
	c.active = false
}

func (c *Customer) SetAddress(address Address) {
	c.address = address
}
