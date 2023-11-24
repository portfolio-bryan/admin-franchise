package postgres

import (
	"context"
	"errors"

	"github.com/bperezgo/admin_franchise/internal/domain/company"
	"github.com/bperezgo/admin_franchise/shared/platform/repositories/postgres"
	"gorm.io/gorm"
)

type CompanyPostgresRepository struct {
	db *gorm.DB
}

func NewCompanyPostgresRepository(db postgres.PostgresRepository) *CompanyPostgresRepository {
	return &CompanyPostgresRepository{
		db: db.PostgresDB,
	}
}

func (c CompanyPostgresRepository) Upsert(ctx context.Context, com company.Company) (company.Company, error) {
	dto := com.DTO()
	comModel := CompanyModel{}

	trx := c.db.First(&comModel, "name = ?",
		dto.Name,
	)

	if errors.Is(trx.Error, gorm.ErrRecordNotFound) {
		trx = c.db.Create(&CompanyModel{
			ID:                dto.ID,
			Name:              dto.Name,
			CompanyOwnerID:    dto.CompanyOwnerID,
			TaxNumber:         dto.TaxNumber,
			AddressLocationID: dto.AddressLocationID,
			LocationID:        dto.LocationID,
		})

		return com, trx.Error
	}

	if trx.Error != nil {
		return company.Company{}, trx.Error
	}

	com, err := company.NewCompany(
		comModel.ID,
		dto.CompanyOwnerID,
		dto.Name,
		dto.TaxNumber,
		dto.LocationID,
		dto.AddressLocationID,
	)
	if err != nil {
		return company.Company{}, err
	}
	// TODO: Create an error for the user, only log the the error
	return com, nil
}

func (c CompanyPostgresRepository) GetByName(ctx context.Context, name string) (company.Company, error) {
	comModel := CompanyModel{}

	trx := c.db.First(&comModel, "name = ?", name)
	if errors.Is(trx.Error, gorm.ErrRecordNotFound) {
		return company.Company{}, company.ErrCompanyNotFound
	}

	if trx.Error != nil {
		return company.Company{}, trx.Error
	}

	com, err := company.NewCompany(
		comModel.ID,
		comModel.CompanyOwnerID,
		comModel.Name,
		comModel.TaxNumber,
		comModel.LocationID,
		comModel.AddressLocationID,
	)
	if err != nil {
		return company.Company{}, err
	}

	return com, nil
}
