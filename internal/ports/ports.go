package ports

import (
	"context"

	"github.com/bperezgo/admin_franchise/internal/domain/company"
	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
	"github.com/bperezgo/admin_franchise/internal/domain/location"
)

type FranchiseRepository interface {
	Upsert(ctx context.Context, franchise franchise.Franchise) error
	SaveIncompleteFranchise(ctx context.Context, franchise franchise.IncompleteFranchise) error

	GetByName(ctx context.Context, name string) (franchise.Franchise, error)
}

type CompanyRepository interface {
	Upsert(ctx context.Context, company company.Company) (company.Company, error)
}

type LocationRepository interface {
	Upsert(ctx context.Context, location location.Location) (location.Location, error)

	UpsertAddress(ctx context.Context, address location.AddressLocation) (location.AddressLocation, error)
}
