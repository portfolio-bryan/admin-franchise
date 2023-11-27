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
	sharedevent "github.com/bperezgo/admin_franchise/shared/domain/event"
	"github.com/google/uuid"
)

type FranchiseCreator struct {
	franchiseRepository ports.FranchiseRepository
	companyRepository   ports.CompanyRepository
	locationRepository  ports.LocationRepository
	eventBus            sharedevent.Bus
}

func NewFranchiseCreator(
	franchiseRepository ports.FranchiseRepository,
	companyRepository ports.CompanyRepository,
	locationRepository ports.LocationRepository,
	eventBus sharedevent.Bus,
) FranchiseCreator {
	return FranchiseCreator{
		franchiseRepository: franchiseRepository,
		companyRepository:   companyRepository,
		locationRepository:  locationRepository,
		eventBus:            eventBus,
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

	locationAggregate, err := location.NewLocation(
		uuid.NewString(),
		scrapResponse.WhoisData.Administrative.Country,
		scrapResponse.WhoisData.Administrative.Province,
		scrapResponse.WhoisData.Administrative.City,
	)
	if err != nil {
		return err
	}

	locationAggregate, err = f.locationRepository.Upsert(ctx, locationAggregate)
	if err != nil {
		return err
	}

	addressLocationAggregate, err := location.NewAddressLocation(
		uuid.NewString(),
		locationAggregate.DTO().ID,
		scrapResponse.WhoisData.Administrative.Street,
		scrapResponse.WhoisData.Administrative.PostalCode,
	)
	if err != nil {
		return err
	}

	addressLocationAggregate, err = f.locationRepository.UpsertAddress(ctx, addressLocationAggregate)
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

	companyAggregate, err = f.companyRepository.Upsert(ctx, companyAggregate)
	if err != nil {
		return err
	}

	franchiseDTO := domainFranchise.FranchiseDTO{
		ID:                   fData.AggregateID,
		URL:                  fData.Url,
		CompanyID:            companyAggregate.DTO().ID,
		Title:                scrapResponse.HTMLMetaData.Title,
		SiteName:             scrapResponse.HTMLMetaData.SiteName,
		Description:          scrapResponse.HTMLMetaData.Description,
		Image:                scrapResponse.HTMLMetaData.Image,
		LocationID:           locationAggregate.DTO().ID,
		AddressLocationID:    addressLocationAggregate.DTO().ID,
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
		// TODO: Publish an event to notify that the franchise is incomplete
	}

	if err := f.franchiseRepository.Upsert(ctx, franchise); err != nil {
		return err
	}

	return f.eventBus.Publish(ctx, franchise.PullEvents())
}
