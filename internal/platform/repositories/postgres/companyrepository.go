package postgres

import (
	"context"

	"github.com/bperezgo/admin_franchise/internal/domain/company"
)

type CompanyPostgresRepository struct {
}

func NewCompanyPostgresRepository() *CompanyPostgresRepository {
	return &CompanyPostgresRepository{}
}

func (c CompanyPostgresRepository) Upsert(ctx context.Context, company company.Company) error {
	return nil
}
