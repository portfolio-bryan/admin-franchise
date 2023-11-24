package postgres

import (
	"context"

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

func (f FranchisePostgresRepository) Save(ctx context.Context, franchise franchise.Franchise) error {
	return nil
}

func (f FranchisePostgresRepository) SaveIncompleteFranchise(ctx context.Context, franchise franchise.IncompleteFranchise) error {
	return nil
}
