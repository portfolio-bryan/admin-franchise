package ports

import (
	"context"

	"github.com/bperezgo/admin_franchise/internal/domain/company"
	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
	"github.com/bperezgo/admin_franchise/internal/domain/location"
)

type FranchiseRepository interface {
	Save(ctx context.Context, franchise franchise.Franchise) error
	SaveIncompleteFranchise(ctx context.Context, franchise franchise.IncompleteFranchise) error
}

type CompanyRepository interface {
	Upsert(ctx context.Context, company company.Company) error
}

type LocationRepository interface {
	Upsert(ctx context.Context, location location.Location) error

	UpsertAddress(ctx context.Context, address location.AddressLocation) error
}
