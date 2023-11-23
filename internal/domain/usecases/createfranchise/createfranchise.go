package createfranchise

import (
	"context"
	"encoding/json"
	"log"

	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
	domainFranchise "github.com/bperezgo/admin_franchise/internal/domain/franchise"
	"github.com/bperezgo/admin_franchise/internal/domain/usecases/scrapfranquise"
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
	franchiseEvt, ok := evt.(franchise.FranchiseRequestReceivedEvent)
	if !ok {
		return isNotFranchiseRequestReceivedEvent
	}

	fData := franchise.FranchiseRequestReceivedEventData{}

	if err := json.Unmarshal(franchiseEvt.Data(), &fData); err != nil {
		return err
	}

	scrapResponse, err := scrapfranquise.ScrapFranquise(ctx, fData.Url)
	if err != nil {
		return err
	}

	log.Println(scrapResponse)

	franchise, err := domainFranchise.NewFranchise("", "", "", "", "", "")
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
