package postgres

import (
	"context"

	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
)

type FranchisePostgresRepository struct {
}

func NewFranchisePostgresRepository() *FranchisePostgresRepository {
	return &FranchisePostgresRepository{}
}

func (f FranchisePostgresRepository) Save(ctx context.Context, franchise franchise.Franchise) error {
	return nil
}
