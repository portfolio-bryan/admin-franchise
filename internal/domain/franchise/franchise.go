package franchise

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/bperezgo/admin_franchise/shared/domain/valueobjects"
	"github.com/google/uuid"
)

type Franchise struct {
	id                FranchiseID
	url               FranchiseURL
	companyId         valueobjects.UID
	name              FranchiseName
	locationId        valueobjects.UID
	addressLocationId valueobjects.UID
}

func NewFranchise(id, url, companyId, name, locationId, addressLocationId string) (Franchise, error) {
	idVO, err := NewFranchiseID(id)
	if err != nil {
		return Franchise{}, err
	}

	urlVO, err := NewFranchiseURL(url)
	if err != nil {
		return Franchise{}, err
	}

	companyIdVO, err := valueobjects.NewUID(companyId)
	if err != nil {
		return Franchise{}, err
	}

	nameVO, err := NewCourseName(name)
	if err != nil {
		return Franchise{}, err
	}

	locationIdVO, err := valueobjects.NewUID(locationId)
	if err != nil {
		return Franchise{}, err
	}

	addressLocationIdVO, err := valueobjects.NewUID(addressLocationId)
	if err != nil {
		return Franchise{}, err
	}

	return Franchise{
		id:                idVO,
		url:               urlVO,
		companyId:         companyIdVO,
		name:              nameVO,
		locationId:        locationIdVO,
		addressLocationId: addressLocationIdVO,
	}, nil
}

var ErrInvalidFranchiseID = errors.New("invalid Course ID")

type FranchiseID struct {
	value string
}

func NewFranchiseID(value string) (FranchiseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return FranchiseID{}, fmt.Errorf("%w: %s", ErrInvalidFranchiseID, value)
	}

	return FranchiseID{
		value: v.String(),
	}, nil
}

type FranchiseURL struct {
	value string
}

var ErrInvalidFranchiseURL = errors.New("invalid Course URL")

func NewFranchiseURL(value string) (FranchiseURL, error) {
	if ok := govalidator.IsURL(value); !ok {
		return FranchiseURL{}, ErrInvalidFranchiseURL
	}

	return FranchiseURL{
		value: value,
	}, nil
}

// FranchiseName
var ErrEmptyFranchiseName = errors.New("the field Course Name can not be empty")

type FranchiseName struct {
	value string
}

func NewCourseName(value string) (FranchiseName, error) {
	if value == "" {
		return FranchiseName{}, ErrEmptyFranchiseName
	}

	return FranchiseName{
		value: value,
	}, nil
}
