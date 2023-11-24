package graph

import (
	"github.com/bperezgo/admin_franchise/config"
	"github.com/bperezgo/admin_franchise/internal/domain/franchise"
	"github.com/bperezgo/admin_franchise/internal/domain/usecases/createfranchise"
	"github.com/bperezgo/admin_franchise/internal/domain/usecases/getcompany"
	"github.com/bperezgo/admin_franchise/internal/domain/usecases/getfranchise"
	repo "github.com/bperezgo/admin_franchise/internal/platform/repositories/postgres"
	"github.com/bperezgo/admin_franchise/shared/platform/event"
	"github.com/bperezgo/admin_franchise/shared/platform/repositories/postgres"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	franchiseCreatorRequestReceiver createfranchise.FranchiseCreatorRequestReceiver
	franchiseGetter                 getfranchise.FranchiseGetter
	companyGetter                   getcompany.CompanyGetter
}

func NewResolver() *Resolver {
	c := config.GetConfig()

	db := postgres.New(postgres.PostgresConfig{
		Host:     c.POSTGRES_HOST,
		Port:     c.POSTGRES_PORT,
		User:     c.POSTGRES_USERNAME,
		Password: c.POSTGRES_PASSWORD,
		DBName:   c.POSTGRES_DATABASE,
	})

	logTrailingDB := event.NewLogTrailingDB(db)

	channelError := event.NewChannelError()
	channelOwner := event.NewChannelOwner(logTrailingDB, channelError)

	// franchiseCreator is an Event Handler
	franchiseRepository := repo.NewFranchisePostgresRepository(db)
	companyRepository := repo.NewCompanyPostgresRepository(db)
	locationRepository := repo.NewLocationPostgresRepository(db)
	franchiseCreator := createfranchise.NewFranchiseCreator(
		franchiseRepository,
		companyRepository,
		locationRepository,
	)
	channelUtilizer := event.NewChannelUtilizer(franchiseCreator, channelError, logTrailingDB)
	channelUtilizer.Use(channelOwner.ChannelEvent(franchise.FranchiseRequestReceivedType))

	franchiseCreatorRequestReceiver := createfranchise.NewFranchiseCreatorRequestReceiver(channelOwner)

	// FranchiseGetter UseCase
	franchiseGetter := getfranchise.NewFranchiseGetter(franchiseRepository)

	// CompanyGetter UseCase
	companyGetter := getcompany.NewCompanyGetter(companyRepository)

	return &Resolver{
		franchiseCreatorRequestReceiver: franchiseCreatorRequestReceiver,
		franchiseGetter:                 franchiseGetter,
		companyGetter:                   companyGetter,
	}
}
