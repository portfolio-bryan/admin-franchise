package getcompany

import (
	"context"

	"github.com/bperezgo/admin_franchise/internal/domain/company"
	"github.com/bperezgo/admin_franchise/internal/ports"
)

type CompanyGetter struct {
	companyRepository ports.CompanyRepository
}

func NewCompanyGetter(companyRepository ports.CompanyRepository) CompanyGetter {
	return CompanyGetter{
		companyRepository: companyRepository,
	}
}

func (g *CompanyGetter) GetCompanyByName(ctx context.Context, name string) (company.Company, error) {
	return g.companyRepository.GetByName(ctx, name)
}
