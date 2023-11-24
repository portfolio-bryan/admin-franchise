package postgres

import "gorm.io/gorm"

type AddressLocationModel struct {
	gorm.Model
	Address string
	ZipCode string
}

type CompanyModel struct {
	gorm.Model
	Name string
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
