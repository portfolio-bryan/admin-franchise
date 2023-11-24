package franchise

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/bperezgo/admin_franchise/shared/domain/valueobjects"
)

type FranchiseDTO struct {
	ID                   string
	URL                  string
	CompanyID            string
	Title                string
	Description          string
	Image                string
	SiteName             string
	Protocol             string
	DomainJumps          int
	ServerNames          []string
	DomainCreationDate   string
	DomainExpirationDate string
	RegistrantName       string
	RegistrantEmail      string
	LocationID           string
	AddressLocationID    string
}

type Franchise struct {
	id                   FranchiseID
	url                  FranchiseURL
	companyId            valueobjects.UID
	title                FranchiseTitle
	description          FranchiseDescription
	image                FranchiseImage
	siteName             FranchiseSiteName
	protocol             Protocol
	domainJumps          FranchiseDomainJumps
	serverNames          FranchiseServerNames
	domainCreationDate   FranchiseDomainCreationDate
	domainExpirationDate FranchiseDomainExpirationDate
	registrantName       FranchiseRegistrantName
	registrantEmail      FranchiseRegistrantEmail
	locationId           valueobjects.UID
	addressLocationId    valueobjects.UID
}

func NewFranchise(franchiseDTO FranchiseDTO) (Franchise, error) {
	idVO, err := NewFranchiseID(franchiseDTO.ID)
	if err != nil {
		return Franchise{}, err
	}

	urlVO, err := NewFranchiseURL(franchiseDTO.URL)
	if err != nil {
		return Franchise{}, err
	}

	companyIdVO, err := valueobjects.NewUID(franchiseDTO.CompanyID)
	if err != nil {
		return Franchise{}, err
	}

	titleVO, err := NewFranchiseTitle(franchiseDTO.Title)
	if err != nil {
		return Franchise{}, err
	}

	descriptionVO, err := NewFranchiseDescription(franchiseDTO.Description)
	if err != nil {
		return Franchise{}, err
	}

	imageVO, err := NewFranchiseImage(franchiseDTO.Image)
	if err != nil {
		return Franchise{}, err
	}

	siteNameVO, err := NewFranchiseSiteName(franchiseDTO.SiteName)
	if err != nil {
		return Franchise{}, err
	}

	protocolVO, err := NewProtocol(franchiseDTO.Protocol)
	if err != nil {
		return Franchise{}, err
	}

	domainJumpsVO, err := NewFranchiseDomainJumps(franchiseDTO.DomainJumps)
	if err != nil {
		return Franchise{}, err
	}

	serverNamesVO, err := NewFranchiseServerNames(franchiseDTO.ServerNames)
	if err != nil {
		return Franchise{}, err
	}

	domainCreationDateVO, err := NewFranchiseDomainCreationDate(franchiseDTO.DomainCreationDate)
	if err != nil {
		return Franchise{}, err
	}

	domainExpirationDateVO, err := NewFranchiseDomainExpirationDate(franchiseDTO.DomainExpirationDate)
	if err != nil {
		return Franchise{}, err
	}

	registrantNameVO, err := NewFranchiseRegistrantName(franchiseDTO.RegistrantName)
	if err != nil {
		return Franchise{}, err
	}

	registrantEmailVO, err := NewFranchiseRegistrantEmail(franchiseDTO.RegistrantEmail)
	if err != nil {
		return Franchise{}, err
	}

	locationIdVO, err := valueobjects.NewUID(franchiseDTO.LocationID)
	if err != nil {
		return Franchise{}, err
	}

	addressLocationIdVO, err := valueobjects.NewUID(franchiseDTO.AddressLocationID)
	if err != nil {
		return Franchise{}, err
	}

	return Franchise{
		id:                   idVO,
		url:                  urlVO,
		companyId:            companyIdVO,
		title:                titleVO,
		description:          descriptionVO,
		image:                imageVO,
		siteName:             siteNameVO,
		protocol:             protocolVO,
		domainJumps:          domainJumpsVO,
		serverNames:          serverNamesVO,
		domainCreationDate:   domainCreationDateVO,
		domainExpirationDate: domainExpirationDateVO,
		registrantName:       registrantNameVO,
		registrantEmail:      registrantEmailVO,
		locationId:           locationIdVO,
		addressLocationId:    addressLocationIdVO,
	}, nil
}

func (f Franchise) DTO() FranchiseDTO {
	return FranchiseDTO{
		ID:                   f.id.value,
		URL:                  f.url.value,
		CompanyID:            f.companyId.String(),
		Title:                f.title.value,
		Description:          f.description.value,
		Image:                f.image.value,
		SiteName:             f.siteName.value,
		Protocol:             f.protocol.value,
		DomainJumps:          f.domainJumps.value,
		ServerNames:          f.serverNames.value,
		DomainCreationDate:   f.domainCreationDate.value,
		DomainExpirationDate: f.domainExpirationDate.value,
		RegistrantName:       f.registrantName.value,
		RegistrantEmail:      f.registrantEmail.value,
		LocationID:           f.locationId.String(),
		AddressLocationID:    f.addressLocationId.String(),
	}
}

var ErrInvalidFranchiseID = errors.New("invalid Franchise ID")

type FranchiseID struct {
	value string
}

func NewFranchiseID(value string) (FranchiseID, error) {
	v, err := valueobjects.NewUID(value)
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

var ErrInvalidFranchiseURL = errors.New("invalid Franchise URL")

func NewFranchiseURL(value string) (FranchiseURL, error) {
	if ok := govalidator.IsURL(value); !ok {
		return FranchiseURL{}, ErrInvalidFranchiseURL
	}

	return FranchiseURL{
		value: value,
	}, nil
}

// FranchiseName
var ErrEmptyFranchiseName = errors.New("the field Franchise Name can not be empty")

type FranchiseTitle struct {
	value string
}

func NewFranchiseTitle(value string) (FranchiseTitle, error) {
	if value == "" {
		return FranchiseTitle{}, ErrEmptyFranchiseName
	}

	return FranchiseTitle{
		value: value,
	}, nil
}

type FranchiseDescription struct {
	value string
}

var ErrEmptyFranchiseDescription = errors.New("the field Franchise Description can not be empty")

func NewFranchiseDescription(value string) (FranchiseDescription, error) {
	if value == "" {
		return FranchiseDescription{}, ErrEmptyFranchiseDescription
	}

	return FranchiseDescription{
		value: value,
	}, nil
}

type FranchiseImage struct {
	value string
}

var ErrInvalidFranchiseImage = errors.New("the field Franchise Image can not be empty")

func NewFranchiseImage(value string) (FranchiseImage, error) {
	if ok := govalidator.IsURL(value); !ok {
		return FranchiseImage{}, ErrInvalidFranchiseImage
	}

	return FranchiseImage{
		value: value,
	}, nil
}

type FranchiseSiteName struct {
	value string
}

var ErrEmptyFranchiseSiteName = errors.New("the field Franchise SiteName can not be empty")

func NewFranchiseSiteName(value string) (FranchiseSiteName, error) {
	if value == "" {
		return FranchiseSiteName{}, ErrEmptyFranchiseSiteName
	}

	return FranchiseSiteName{
		value: value,
	}, nil
}

type Protocol struct {
	value string
}

var ErrInvalidProtocol = errors.New("invalid Protocol")

func NewProtocol(value string) (Protocol, error) {
	// Provisional logic
	if value != "http" && value != "https" {
		return Protocol{}, fmt.Errorf("%w: %s", ErrInvalidProtocol, value)
	}

	return Protocol{
		value: value,
	}, nil
}

type FranchiseDomainJumps struct {
	value int
}

var ErrInvalidFranchiseDomainJumps = errors.New("invalid Franchise Domain Jumps")

func NewFranchiseDomainJumps(value int) (FranchiseDomainJumps, error) {
	if value < 0 {
		return FranchiseDomainJumps{}, ErrInvalidFranchiseDomainJumps
	}

	return FranchiseDomainJumps{
		value: value,
	}, nil
}

type FranchiseServerNames struct {
	value []string
}

var ErrInvalidFranchiseServerNames = errors.New("invalid Franchise Server Names")

func NewFranchiseServerNames(value []string) (FranchiseServerNames, error) {
	return FranchiseServerNames{
		value: value,
	}, nil
}

type FranchiseDomainCreationDate struct {
	value string
}

func NewFranchiseDomainCreationDate(value string) (FranchiseDomainCreationDate, error) {
	return FranchiseDomainCreationDate{
		value: value,
	}, nil
}

type FranchiseDomainExpirationDate struct {
	value string
}

func NewFranchiseDomainExpirationDate(value string) (FranchiseDomainExpirationDate, error) {
	return FranchiseDomainExpirationDate{
		value: value,
	}, nil
}

type FranchiseRegistrantName struct {
	value string
}

var ErrEmptyFranchiseRegistrantName = errors.New("the field Franchise Registrant Name can not be empty")

func NewFranchiseRegistrantName(value string) (FranchiseRegistrantName, error) {
	if value == "" {
		return FranchiseRegistrantName{}, ErrEmptyFranchiseRegistrantName
	}

	return FranchiseRegistrantName{
		value: value,
	}, nil
}

type FranchiseRegistrantEmail struct {
	value string
}

var ErrInvalidFranchiseRegistrantEmail = errors.New("invalid Franchise Registrant Email")

func NewFranchiseRegistrantEmail(value string) (FranchiseRegistrantEmail, error) {
	if ok := govalidator.IsEmail(value); !ok {
		return FranchiseRegistrantEmail{}, ErrInvalidFranchiseRegistrantEmail
	}

	return FranchiseRegistrantEmail{
		value: value,
	}, nil
}
