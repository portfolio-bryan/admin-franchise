package postgres

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
	"github.com/bperezgo/admin_franchise/internal/domain/views"
	"github.com/bperezgo/admin_franchise/shared/platform/repositories/postgres"
	"gorm.io/gorm"
)

type FranchisePostgresRepository struct {
	db *gorm.DB
}

func NewFranchisePostgresRepository(db postgres.PostgresRepository) *FranchisePostgresRepository {
	return &FranchisePostgresRepository{
		db: db.PostgresDB,
	}
}

func (f FranchisePostgresRepository) Upsert(ctx context.Context, fran franchise.Franchise) error {
	dto := fran.DTO()
	model := FranchiseModel{}

	trx := f.db.First(&model, "title = ? AND company_id = ?",
		dto.Title,
		dto.CompanyID,
	)

	if errors.Is(trx.Error, gorm.ErrRecordNotFound) {
		trx = f.db.Create(&FranchiseModel{
			ID:                   dto.ID,
			CompanyID:            dto.CompanyID,
			Title:                dto.Title,
			SiteName:             dto.SiteName,
			Description:          dto.Description,
			Image:                dto.Image,
			URL:                  dto.URL,
			Protocol:             dto.Protocol,
			DomainJumps:          dto.DomainJumps,
			ServerNames:          dto.ServerNames,
			DomainCreationDate:   dto.DomainCreationDate,
			DomainExpirationDate: dto.DomainExpirationDate,
			RegistrantName:       dto.RegistrantName,
			RegistrantEmail:      dto.RegistrantEmail,
			LocationID:           dto.LocationID,
			AddressLocationID:    dto.AddressLocationID,
		})

		return trx.Error
	}

	if trx.Error != nil {
		return trx.Error
	}

	// TODO: Create an error for the user, only log the the error
	return nil
}

func (f FranchisePostgresRepository) SaveIncompleteFranchise(ctx context.Context, fran franchise.IncompleteFranchise) error {
	dto := fran.Data

	b, err := json.Marshal(&dto)

	if err != nil {
		return err
	}

	trx := f.db.Create(&IncompleteFranchiseModel{
		ID:                dto.ID,
		Data:              string(b),
		WasVerified:       false,
		URL:               dto.URL,
		Name:              dto.Title,
		LocationID:        dto.LocationID,
		AddressLocationID: dto.AddressLocationID,
	})

	// TODO: Create an error for the user, only log the the error
	return trx.Error
}

func (f FranchisePostgresRepository) GetByName(ctx context.Context, name string) (views.Franchise, error) {
	model := FranchiseModel{}

	trx := f.db.Joins("Company").Joins("Location").Joins("AddressLocation").Where("site_name = ?", name).First(&model)

	if errors.Is(trx.Error, gorm.ErrRecordNotFound) {
		return views.Franchise{}, franchise.ErrFranchiseNotFound
	}

	if trx.Error != nil {
		return views.Franchise{}, trx.Error
	}

	return views.Franchise{
		ID:                   model.ID,
		URL:                  model.URL,
		CompanyID:            model.CompanyID,
		Title:                model.Title,
		Description:          model.Description,
		Image:                model.Image,
		SiteName:             model.SiteName,
		Protocol:             model.Protocol,
		DomainJumps:          model.DomainJumps,
		ServerNames:          model.ServerNames,
		DomainCreationDate:   model.DomainCreationDate,
		DomainExpirationDate: model.DomainExpirationDate,
		RegistrantName:       model.RegistrantName,
		RegistrantEmail:      model.RegistrantEmail,
		LocationID:           model.LocationID,
		AddressLocationID:    model.AddressLocationID,
		Company: views.Company{
			ID:                model.Company.ID,
			Name:              model.Company.Name,
			CompanyOwnerID:    model.Company.CompanyOwnerID,
			TaxNumber:         model.Company.TaxNumber,
			LocationID:        model.Company.LocationID,
			AddressLocationID: model.Company.AddressLocationID,
		},
		Location: views.Location{
			ID:      model.Location.ID,
			City:    model.Location.City,
			Country: model.Location.Country,
			State:   model.Location.State,
		},
		AddressLocation: views.AddressLocation{
			ID:         model.AddressLocation.ID,
			LocationID: model.AddressLocation.LocationID,
			Address:    model.AddressLocation.Address,
			ZipCode:    model.AddressLocation.ZipCode,
		},
	}, nil
}
