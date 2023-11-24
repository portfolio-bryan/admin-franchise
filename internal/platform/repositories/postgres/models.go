package postgres

import "gorm.io/gorm"

type AddressLocationModel struct {
	ID         string
	LocationID string
	Address    string
	ZipCode    string
}

func (AddressLocationModel) TableName() string {
	return "address_location"
}

type CompanyModel struct {
	gorm.Model
	ID                string
	Name              string
	CompanyOwnerID    string
	TaxNumber         string
	LocationID        string
	AddressLocationID string
}

type FranchiseModel struct {
	gorm.Model
	Name string
}

type CompanyOwnerModel struct {
	gorm.Model
	CompanyID uint
	Id        uint
}

type LocationModel struct {
	ID      string
	City    string
	Country string
	State   string
}

func (LocationModel) TableName() string {
	return "locations"
}
