package location

import (
	"errors"

	"github.com/bperezgo/admin_franchise/shared/domain/valueobjects"
)

type AddressLocationDTO struct {
	ID string
}

type AddressLocation struct {
	ID         AddressLocationID
	LocationID valueobjects.UID
	Address    Address
	ZipCode    ZipCode
}

func NewAddressLocation(
	id string,
	locationID string,
	address string,
	zipCode string,
) (AddressLocation, error) {
	idVO, err := NewAddressLocationID(id)
	if err != nil {
		return AddressLocation{}, err
	}

	locationIDVO, err := valueobjects.NewUID(locationID)
	if err != nil {
		return AddressLocation{}, err
	}

	addressVO, err := NewAddress(address)
	if err != nil {
		return AddressLocation{}, err
	}

	zipCodeVO, err := NewZipCode(zipCode)
	if err != nil {
		return AddressLocation{}, err
	}

	return AddressLocation{
		ID:         idVO,
		LocationID: locationIDVO,
		Address:    addressVO,
		ZipCode:    zipCodeVO,
	}, nil
}

func (a AddressLocation) DTO() AddressLocationDTO {
	return AddressLocationDTO{
		ID: a.ID.value,
	}
}

type AddressLocationID struct {
	value string
}

func NewAddressLocationID(value string) (AddressLocationID, error) {
	v, err := valueobjects.NewUID(value)
	if err != nil {
		return AddressLocationID{}, err
	}

	return AddressLocationID{
		value: v.String(),
	}, nil
}

type Address struct {
	value string
}

var ErrInvalidAddress = errors.New("invalid Address")

func NewAddress(value string) (Address, error) {
	// Provisional logic
	if value == "" {
		return Address{}, ErrInvalidAddress
	}

	return Address{
		value: value,
	}, nil
}

type ZipCode struct {
	value string
}

var ErrInvalidZipCode = errors.New("invalid ZipCode")

func NewZipCode(value string) (ZipCode, error) {
	if value == "" {
		return ZipCode{}, ErrInvalidZipCode
	}

	return ZipCode{
		value: value,
	}, nil
}
