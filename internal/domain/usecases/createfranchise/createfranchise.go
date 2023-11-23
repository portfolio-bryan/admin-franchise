package createfranchise

import (
	"context"

	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
	domainFranchise "github.com/bperezgo/admin_franchise/internal/domain/franchise"
	"github.com/bperezgo/admin_franchise/internal/ports"
	"github.com/bperezgo/admin_franchise/shared/domain/event"
)

type FranchiseCreator struct {
	franchiseRepository ports.FranchiseRepository
}

func NewFranchiseCreator(franchiseRepository ports.FranchiseRepository) FranchiseCreator {
	return FranchiseCreator{
		franchiseRepository: franchiseRepository,
	}
}

func (f FranchiseCreator) Handle(ctx context.Context, evt event.Event) error {
	// Scrap the web to get the data to build the franchise

	_, ok := evt.(franchise.FranchiseRequestReceivedEvent)
	if !ok {
		return isNotFranchiseRequestReceivedEvent
	}

	franchise, err := domainFranchise.NewFranchise()
	if err != nil {
		return f.CreateIncompleteFranchise(ctx, domainFranchise.NewIncompleteFranchise())
	}

	return f.Create(ctx, franchise)
}

func (f FranchiseCreator) Create(ctx context.Context, franchise domainFranchise.Franchise) error {
	return f.franchiseRepository.Save(ctx, franchise)
}

func (f FranchiseCreator) CreateIncompleteFranchise(ctx context.Context, franchise domainFranchise.IncompleteFranchise) error {
	return f.franchiseRepository.SaveIncompleteFranchise(ctx, franchise)
}
