package postgres

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
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
	model := FranchiseModel{
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
	}

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

		return nil
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
		ID:          dto.ID,
		Data:        string(b),
		WasVerified: false,
		URL:         dto.URL,
		Name:        dto.Title,
	})

	// TODO: Create an error for the user, only log the the error
	return trx.Error
}
