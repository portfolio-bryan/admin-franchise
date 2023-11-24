package postgres

import (
	"context"

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

func (c CompanyPostgresRepository) Upsert(ctx context.Context, company company.Company) error {
	return nil
}
