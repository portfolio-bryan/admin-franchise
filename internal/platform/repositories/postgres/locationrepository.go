package postgres

import (
	"context"

	"github.com/bperezgo/admin_franchise/internal/domain/location"
)

type LocationPostgresRepository struct {
}

func NewLocationPostgresRepository() *LocationPostgresRepository {
	return &LocationPostgresRepository{}
}

func (l LocationPostgresRepository) Upsert(ctx context.Context, location location.Location) error {
	return nil
}

func (l LocationPostgresRepository) UpsertAddress(ctx context.Context, address location.AddressLocation) error {
	return nil
}
