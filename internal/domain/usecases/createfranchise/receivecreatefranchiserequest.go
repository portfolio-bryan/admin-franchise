package createfranchise

import (
	"context"

	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
	sharedevent "github.com/bperezgo/admin_franchise/shared/domain/event"
	"github.com/go-playground/validator/v10"
)

type FranchiseCreatorRequestReceiver struct {
	eventBus sharedevent.Bus
}

func NewFranchiseCreatorRequestReceiver(eventBus sharedevent.Bus) FranchiseCreatorRequestReceiver {
	return FranchiseCreatorRequestReceiver{
		eventBus: eventBus,
	}
}

func (f FranchiseCreatorRequestReceiver) Receive(ctx context.Context, createDTO franchise.CreateDTO) error {
	// Validate the DTO
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(createDTO)
	if err != nil {
		return err
	}

	event := franchise.NewFranchiseRequestReceivedEvent(createDTO.ID, createDTO.URL)
	if err := f.eventBus.Publish(ctx, []sharedevent.Event{event}); err != nil {
		return err
	}

	return nil
}
