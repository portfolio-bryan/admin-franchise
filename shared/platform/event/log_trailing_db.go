package event

import (
	"github.com/bperezgo/admin_franchise/shared/domain/event"
	"github.com/bperezgo/admin_franchise/shared/platform/repositories/postgres"
	"gorm.io/gorm"
)

type LogTrailingDB struct {
	db *gorm.DB
}

func NewLogTrailingDB(db postgres.PostgresRepository) LogTrailingDB {
	return LogTrailingDB{
		db: db.PostgresDB,
	}
}

func (l LogTrailingDB) SavePendingEvent(evt event.Event) error {
	trx := l.db.Create(&TransactionLogTrailing{
		EventID:    evt.ID(),
		Data:       string(evt.Data()),
		WasRead:    false,
		EventType:  string(evt.Type()),
		OccurredOn: evt.OccurredOn(),
	})

	// TODO: Create an error for the user, only log the the error
	return trx.Error
}

func (l LogTrailingDB) FulfillEvent(evt event.Event) error {
	trx := l.db.Where("event_id = ?", evt.ID()).Save(&TransactionLogTrailing{
		EventID:    evt.ID(),
		Data:       string(evt.Data()),
		WasRead:    true,
		EventType:  string(evt.Type()),
		OccurredOn: evt.OccurredOn(),
	})

	// TODO: Create an error for the user, only log the the error
	return trx.Error
}
