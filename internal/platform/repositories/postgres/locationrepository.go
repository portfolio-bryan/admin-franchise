package postgres

import (
	"context"
	"errors"

	"github.com/bperezgo/admin_franchise/internal/domain/location"
	"github.com/bperezgo/admin_franchise/shared/platform/repositories/postgres"
	"gorm.io/gorm"
)

type LocationPostgresRepository struct {
	db *gorm.DB
}

func NewLocationPostgresRepository(db postgres.PostgresRepository) *LocationPostgresRepository {
	return &LocationPostgresRepository{
		db: db.PostgresDB,
	}
}

func (l LocationPostgresRepository) Upsert(ctx context.Context, loc location.Location) (location.Location, error) {
	dto := loc.DTO()
	locationModel := LocationModel{
		City:    dto.City,
		Country: dto.Country,
		State:   dto.State,
	}

	trx := l.db.First(&locationModel, "city = ? AND country = ? AND state = ?",
		dto.City,
		dto.Country,
		dto.State,
	)

	if errors.Is(trx.Error, gorm.ErrRecordNotFound) {
		trx = l.db.Create(&LocationModel{
			ID:      dto.ID,
			City:    dto.City,
			Country: dto.Country,
			State:   dto.State,
		})

		return loc, nil
	}

	loc, err := location.NewLocation(locationModel.ID, dto.Country, dto.State, dto.City)
	if err != nil {
		return location.Location{}, err
	}
	// TODO: Create an error for the user, only log the the error
	return loc, nil
}

func (l LocationPostgresRepository) UpsertAddress(ctx context.Context, addLoc location.AddressLocation) (location.AddressLocation, error) {
	dto := addLoc.DTO()
	addLocationModel := AddressLocationModel{
		ID:         dto.ID,
		LocationID: dto.LocationID,
		Address:    dto.Address,
		ZipCode:    dto.ZipCode,
	}

	trx := l.db.First(&addLocationModel, "location_id = ? AND address = ? AND zip_code = ?",
		dto.LocationID,
		dto.Address,
		dto.ZipCode,
	)

	if errors.Is(trx.Error, gorm.ErrRecordNotFound) {
		trx = l.db.Create(&AddressLocationModel{
			ID:         dto.ID,
			LocationID: dto.LocationID,
			Address:    dto.Address,
			ZipCode:    dto.ZipCode,
		})

		return addLoc, nil
	}

	loc, err := location.NewAddressLocation(addLocationModel.ID, dto.LocationID, dto.Address, dto.ZipCode)
	if err != nil {
		return location.AddressLocation{}, err
	}
	// TODO: Create an error for the user, only log the the error
	return loc, nil
}
