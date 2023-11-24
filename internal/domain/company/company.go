package company

import "github.com/bperezgo/admin_franchise/shared/domain/valueobjects"

type CompanyDTO struct {
	ID                string
	CompanyOwnerID    string
	Name              string
	TaxNumber         string
	LocationID        string
	AddressLocationID string
}

type Company struct {
	ID                CompanyID
	CompanyOwnerID    valueobjects.UID
	Name              CompanyName
	TaxNumber         CompanyTaxNumber
	LocationID        valueobjects.UID
	AddressLocationID valueobjects.UID
}

func NewCompany(id, companyOwnerID, name, taxNumber, locationID, addressLocationID string) (Company, error) {
	idVO, err := NewCompanyID(id)
	if err != nil {
		return Company{}, err
	}

	coVO, err := valueobjects.NewUID(companyOwnerID)
	if err != nil {
		return Company{}, err
	}

	nameVO, err := NewCompanyName(name)
	if err != nil {
		return Company{}, err
	}

	taxNumberVO, err := NewCompanyTaxNumber(taxNumber)
	if err != nil {
		return Company{}, err
	}

	locationIDVO, err := valueobjects.NewUID(locationID)
	if err != nil {
		return Company{}, err
	}

	addressLocationIDVO, err := valueobjects.NewUID(addressLocationID)
	if err != nil {
		return Company{}, err
	}

	return Company{
		ID:                idVO,
		CompanyOwnerID:    coVO,
		Name:              nameVO,
		TaxNumber:         taxNumberVO,
		LocationID:        locationIDVO,
		AddressLocationID: addressLocationIDVO,
	}, nil
}

func (c Company) DTO() CompanyDTO {
	return CompanyDTO{
		ID:                c.ID.value,
		CompanyOwnerID:    c.CompanyOwnerID.String(),
		Name:              c.Name.Value,
		TaxNumber:         c.TaxNumber.Value,
		LocationID:        c.LocationID.String(),
		AddressLocationID: c.AddressLocationID.String(),
	}
}

type CompanyID struct {
	value string
}

func NewCompanyID(value string) (CompanyID, error) {
	v, err := valueobjects.NewUID(value)
	if err != nil {
		return CompanyID{}, err
	}

	return CompanyID{
		value: v.String(),
	}, nil
}

type CompanyName struct {
	Value string
}

func NewCompanyName(value string) (CompanyName, error) {
	return CompanyName{
		Value: value,
	}, nil
}

type CompanyTaxNumber struct {
	Value string
}

func NewCompanyTaxNumber(value string) (CompanyTaxNumber, error) {
	return CompanyTaxNumber{
		Value: value,
	}, nil
}
