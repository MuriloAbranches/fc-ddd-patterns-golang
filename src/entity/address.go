package entity

import (
	"errors"
	"fmt"
)

var (
	ErrorStreetIsRequired = errors.New("street is required")
	ErrorNumberIsRequired = errors.New("number is required")
	ErrorZipIsRequired    = errors.New("zip is required")
	ErrorCityIsRequired   = errors.New("city is required")
)

type Address struct {
	street string
	number int64
	zip    string
	city   string
}

func NewAddress(street string, number int64, zip, city string) (*Address, error) {
	address := &Address{
		street: street,
		number: number,
		zip:    zip,
		city:   city,
	}

	err := address.Validate()
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (a *Address) Validate() error {
	if a.street == "" {
		return ErrorStreetIsRequired
	}

	if a.number == 0 {
		return ErrorNumberIsRequired
	}

	if a.zip == "" {
		return ErrorZipIsRequired
	}

	if a.street == "" {
		return ErrorCityIsRequired
	}

	return nil
}

func (a *Address) ToString() string {
	return fmt.Sprintf("%s, %d, %s, %s", a.street, a.number, a.zip, a.city)
}
