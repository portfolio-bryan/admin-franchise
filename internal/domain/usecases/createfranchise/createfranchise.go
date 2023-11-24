package createfranchise

import (
	"context"
	"encoding/json"

	"github.com/bperezgo/admin_franchise/internal/domain/company"
	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
	domainFranchise "github.com/bperezgo/admin_franchise/internal/domain/franchise"
	"github.com/bperezgo/admin_franchise/internal/domain/location"
	"github.com/bperezgo/admin_franchise/internal/domain/usecases/scrapfranquise"
	"github.com/bperezgo/admin_franchise/internal/ports"
	"github.com/bperezgo/admin_franchise/shared/domain/event"
	"github.com/google/uuid"
)

type FranchiseCreator struct {
	franchiseRepository ports.FranchiseRepository
	companyRepository   ports.CompanyRepository
	locationRepository  ports.LocationRepository
}

func NewFranchiseCreator(
	franchiseRepository ports.FranchiseRepository,
	companyRepository ports.CompanyRepository,
	locationRepository ports.LocationRepository,
) FranchiseCreator {
	return FranchiseCreator{
		franchiseRepository: franchiseRepository,
		companyRepository:   companyRepository,
		locationRepository:  locationRepository,
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

	locationAggregate, err := location.NewLocation(uuid.NewString(), "country", "state", "city")
	if err != nil {
		return err
	}

	err = f.locationRepository.Upsert(ctx, locationAggregate)
	if err != nil {
		return err
	}

	addressLocationAggregate, err := location.NewAddressLocation(uuid.NewString(), locationAggregate.DTO().ID, "address", "zipCode")
	if err != nil {
		return err
	}

	err = f.locationRepository.UpsertAddress(ctx, addressLocationAggregate)
	if err != nil {
		return err
	}

	companyAggregate, err := company.NewCompany(
		uuid.NewString(),
		// TODO: get company owner id
		uuid.NewString(),
		scrapResponse.WhoisData.Registrant.Name,
		scrapResponse.WhoisData.Administrative.Fax,
		locationAggregate.DTO().ID,
		addressLocationAggregate.DTO().ID,
	)
	if err != nil {
		return err
	}

	err = f.companyRepository.Upsert(ctx, companyAggregate)
	if err != nil {
		return err
	}

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
		return f.franchiseRepository.SaveIncompleteFranchise(ctx, domainFranchise.NewIncompleteFranchise(franchiseDTO))
	}

	return f.franchiseRepository.Save(ctx, franchise)
}
