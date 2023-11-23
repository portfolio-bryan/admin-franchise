package postgres

import "gorm.io/gorm"

type AddressLocation struct {
	gorm.Model
	Address string
	ZipCode string
}

type Company struct {
	gorm.Model
	Name string
}

type Franchise struct {
	gorm.Model
	Name string
}

type CompanyOwner struct {
	gorm.Model
	CompanyID uint
	Id        uint
}

type Location struct {
	gorm.Model
	City    string
	Country string
	State   string
}
