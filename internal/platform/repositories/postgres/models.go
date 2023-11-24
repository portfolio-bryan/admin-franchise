package postgres

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type AddressLocationModel struct {
	ID         string
	LocationID string
	Address    string
	ZipCode    string

	Location LocationModel `gorm:"foreignKey:LocationID;references:ID"`
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

	// Franchises      []FranchiseModel
	Location        LocationModel        `gorm:"foreignKey:LocationID;references:ID"`
	AddressLocation AddressLocationModel `gorm:"foreignKey:AddressLocationID;references:ID"`
}

func (CompanyModel) TableName() string {
	return "company"
}

type FranchiseModel struct {
	gorm.Model
	ID                   string
	CompanyID            string
	Title                string
	SiteName             string
	Description          string
	Image                string
	URL                  string
	Protocol             string
	DomainJumps          int
	ServerNames          pq.StringArray `gorm:"type:text[]"`
	DomainCreationDate   string
	DomainExpirationDate string
	RegistrantName       string
	RegistrantEmail      string
	LocationID           string
	AddressLocationID    string

	Company         CompanyModel         `gorm:"foreignKey:CompanyID;references:ID"`
	Location        LocationModel        `gorm:"foreignKey:LocationID;references:ID"`
	AddressLocation AddressLocationModel `gorm:"foreignKey:AddressLocationID;references:ID"`
}

func (FranchiseModel) TableName() string {
	return "franchise"
}

type IncompleteFranchiseModel struct {
	gorm.Model
	ID                string
	Data              string
	WasVerified       bool
	URL               string
	Name              string
	LocationID        string
	AddressLocationID string
}

func (IncompleteFranchiseModel) TableName() string {
	return "incomplete_franchise"
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
