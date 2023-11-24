package utilstests

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bperezgo/admin_franchise/config"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func UpMigrations(host, port string) {
	log.Println("migration up started")

	var migration = buildMigrate(host, port)

	migrationError := migration.Up()

	if migrationError != nil {
		log.Fatalln("fail when migration up execution : ", migrationError.Error())
	}

	log.Println("migration up finished")
}

func DownMigrations(host, port string) {
	log.Println("migration down started")

	var migration = buildMigrate(host, port)

	migrationError := migration.Down()

	if migrationError != nil {
		log.Fatalln("fail when migration down execution : ", migrationError.Error())
	}

	log.Println("migration down finished")
}

func buildMigrate(host, port string) *migrate.Migrate {
	c := config.GetConfig()

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			c.POSTGRES_USERNAME,
			c.POSTGRES_PASSWORD,
			host,
			port,
			c.POSTGRES_DATABASE,
		),
	)
	if err != nil {
		log.Fatalln("Failing opening DB : ", err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalln("Failing getting the instance : ", err.Error())
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://../../setup/migrations",
		"admin_franchise", driver)
	if err != nil {
		log.Fatalln("error creating new migrate instance : ", err.Error())
	}

	return migration
}
