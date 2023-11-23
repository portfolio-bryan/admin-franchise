package createfranchise

import (
	"context"

	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
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
		// TODO: return error
		return nil
	}

	franchise, err := franchise.NewFranchise()
	if err != nil {
		return err
	}

	return f.Create(ctx, franchise)
}

func (f FranchiseCreator) Create(ctx context.Context, franchise franchise.Franchise) error {

	return f.franchiseRepository.Save(ctx, franchise)
}
