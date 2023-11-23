package graph

import (
	"github.com/bperezgo/admin_franchise/internal/domain/usecases/createfranchise"
	"github.com/bperezgo/admin_franchise/internal/platform/repositories/postgres"
	"github.com/bperezgo/admin_franchise/shared/platform/event"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	franchiseCreatorRequestReceiver createfranchise.FranchiseCreatorRequestReceiver
}

func NewResolver() *Resolver {

	channelOwner := event.NewChannelOwner()
	channelError := event.NewChannelError()

	// franchiseCreator is an Event Handler
	franchiseRepository := postgres.NewFranchisePostgresRepository()
	franchiseCreator := createfranchise.NewFranchiseCreator(franchiseRepository)
	channelUtilizer := event.NewChannelUtilizer(franchiseCreator, channelError)
	channelUtilizer.Use(channelOwner.ChannelEvents())

	franchiseCreatorRequestReceiver := createfranchise.NewFranchiseCreatorRequestReceiver(channelOwner)
	return &Resolver{
		franchiseCreatorRequestReceiver: franchiseCreatorRequestReceiver,
	}
}
