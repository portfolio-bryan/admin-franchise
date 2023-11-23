package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

// TODO: env from config struct
func Init() {
	// Getting postgres driver
	dsn := "host=localhost user=user password=password dbname=admin_franchise port=5432 sslmode=disable TimeZone=america/bogota"
	dbInitiated, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	PostgresDB = dbInitiated

	// TODO: Migrate the schema
}
