package createfranchise

import (
	"context"
	"encoding/json"

	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
	domainFranchise "github.com/bperezgo/admin_franchise/internal/domain/franchise"
	"github.com/bperezgo/admin_franchise/internal/domain/usecases/scrapfranquise"
	"github.com/bperezgo/admin_franchise/internal/ports"
	"github.com/bperezgo/admin_franchise/shared/domain/event"
)

type FranchiseCreator struct {
	franchiseRepository ports.FranchiseRepository
}

func NewFranchiseCreator(franchiseRepository ports.FranchiseRepository) FranchiseCreator {
	return FranchiseCreator{
		franchiseRepository: franchiseRepository,
	}
}

func (f FranchiseCreator) Handle(ctx context.Context, evt event.Event) error {
	franchiseEvt, ok := evt.(franchise.FranchiseRequestReceivedEvent)
	if !ok {
		return isNotFranchiseRequestReceivedEvent
	}

	fData := franchise.FranchiseRequestReceivedEventData{}

	if err := json.Unmarshal(franchiseEvt.Data(), &fData); err != nil {
		return err
	}

	scrapResponse, err := scrapfranquise.ScrapFranquise(ctx, fData.Url)
	if err != nil {
		return err
	}

	// CreateCompany
	// CreateLocation
	// CreateAddressLocation

	franchiseDTO := domainFranchise.FranchiseDTO{
		ID:                   fData.AggregateID,
		URL:                  fData.Url,
		CompanyID:            "companyID",
		Title:                scrapResponse.HTMLMetaData.Title,
		SiteName:             scrapResponse.HTMLMetaData.SiteName,
		Description:          scrapResponse.HTMLMetaData.Description,
		Image:                scrapResponse.HTMLMetaData.Image,
		LocationID:           "locationID",
		AddressLocationID:    "addressLocationID",
		Protocol:             scrapResponse.Protocol,
		DomainJumps:          scrapResponse.Jumps,
		ServerNames:          scrapResponse.WhoisData.Domain.NameServers,
		DomainCreationDate:   scrapResponse.WhoisData.Domain.CreatedDate,
		DomainExpirationDate: scrapResponse.WhoisData.Domain.ExpirationDate,
		RegistrantName:       scrapResponse.WhoisData.Registrant.Name,
		RegistrantEmail:      scrapResponse.WhoisData.Registrant.Email,
	}

	franchise, err := domainFranchise.NewFranchise(franchiseDTO)

	if err != nil {
		return f.CreateIncompleteFranchise(ctx, domainFranchise.NewIncompleteFranchise(franchiseDTO))
	}

	return f.Create(ctx, franchise)
}

func (f FranchiseCreator) Create(ctx context.Context, franchise domainFranchise.Franchise) error {
	return f.franchiseRepository.Save(ctx, franchise)
}

func (f FranchiseCreator) CreateIncompleteFranchise(ctx context.Context, franchise domainFranchise.IncompleteFranchise) error {
	return f.franchiseRepository.SaveIncompleteFranchise(ctx, franchise)
}
