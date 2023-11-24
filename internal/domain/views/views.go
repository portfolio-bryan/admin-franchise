package views

type Location struct {
	ID      string
	City    string
	Country string
	State   string
}

type AddressLocation struct {
	ID         string
	LocationID string
	Address    string
	ZipCode    string
}

type Company struct {
	ID                string
	Name              string
	CompanyOwnerID    string
	TaxNumber         string
	LocationID        string
	AddressLocationID string
}

type Franchise struct {
	ID                   string
	CompanyID            string
	Title                string
	SiteName             string
	Description          string
	Image                string
	URL                  string
	Protocol             string
	DomainJumps          int
	ServerNames          []string
	DomainCreationDate   string
	DomainExpirationDate string
	RegistrantName       string
	RegistrantEmail      string
	LocationID           string
	AddressLocationID    string

	Company         Company
	Location        Location
	AddressLocation AddressLocation
}
