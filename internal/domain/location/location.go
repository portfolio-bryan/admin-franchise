package location

import (
	"errors"

	"github.com/bperezgo/admin_franchise/shared/domain/valueobjects"
)

type LocationDTO struct {
	ID      string
	Country string
	State   string
	City    string
}

type Location struct {
	ID      LocationID
	Country Country
	State   State
	City    City
}

func NewLocation(id, country, state, city string) (Location, error) {
	idVO, err := NewLocationID(id)
	if err != nil {
		return Location{}, err
	}

	countryVO, err := NewCountry(country)
	if err != nil {
		return Location{}, err
	}

	stateVO, err := NewState(state)
	if err != nil {
		return Location{}, err
	}

	cityVO, err := NewCity(city)
	if err != nil {
		return Location{}, err
	}

	return Location{
		ID:      idVO,
		Country: countryVO,
		State:   stateVO,
		City:    cityVO,
	}, nil
}

func (l Location) DTO() LocationDTO {
	return LocationDTO{
		ID: l.ID.Value,
	}
}

// VALUE OBJECTS
type LocationID struct {
	Value string
}

func NewLocationID(value string) (LocationID, error) {
	v, err := valueobjects.NewUID(value)
	if err != nil {
		return LocationID{}, err
	}

	return LocationID{
		Value: v.String(),
	}, nil
}

type Country struct {
	value string
}

var ErrInvalidCountry = errors.New("invalid Country")

func NewCountry(value string) (Country, error) {
	if value == "" {
		return Country{}, ErrInvalidCountry
	}

	return Country{
		value: value,
	}, nil
}

type State struct {
	value string
}

var ErrInvalidState = errors.New("invalid State")

func NewState(value string) (State, error) {
	if value == "" {
		return State{}, ErrInvalidState
	}

	return State{
		value: value,
	}, nil
}

type City struct {
	value string
}

var ErrInvalidCity = errors.New("invalid City")

func NewCity(value string) (City, error) {
	if value == "" {
		return City{}, ErrInvalidCity
	}

	return City{
		value: value,
	}, nil
}
