package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type PostgresRepository struct {
	PostgresDB *gorm.DB
}

func New(config PostgresConfig) PostgresRepository {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=america/bogota",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return PostgresRepository{
		PostgresDB: db,
	}
}
