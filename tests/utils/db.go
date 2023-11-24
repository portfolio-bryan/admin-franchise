package utilstests

import (
	"log"

	"github.com/bperezgo/admin_franchise/config"
	"github.com/bperezgo/admin_franchise/shared/platform/repositories/postgres"
)

var postgresRepository *postgres.PostgresRepository

func Connect(host, port string) *postgres.PostgresRepository {
	if postgresRepository != nil {
		return postgresRepository
	}

	c := config.GetConfig()

	pr := postgres.New(postgres.PostgresConfig{
		Host:     host,
		Port:     port,
		User:     c.POSTGRES_USERNAME,
		Password: c.POSTGRES_PASSWORD,
		DBName:   c.POSTGRES_DATABASE,
	})

	log.Println("Database connection started")

	postgresRepository = &pr

	return postgresRepository
}
