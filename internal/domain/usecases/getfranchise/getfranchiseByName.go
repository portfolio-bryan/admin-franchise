package getfranchise

import (
	"context"

	"github.com/bperezgo/admin_franchise/internal/domain/views"
	"github.com/bperezgo/admin_franchise/internal/ports"
)

type FranchiseGetter struct {
	franchiseRepository ports.FranchiseRepository
}

func NewFranchiseGetter(franchiseRepository ports.FranchiseRepository) FranchiseGetter {
	return FranchiseGetter{
		franchiseRepository: franchiseRepository,
	}
}

func (f FranchiseGetter) GetFranchiseByName(ctx context.Context, name string) (views.Franchise, error) {
	return f.franchiseRepository.GetByName(ctx, name)
}
