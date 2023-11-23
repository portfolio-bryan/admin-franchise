package event

import (
	"time"
)

type TransactionLogTrailing struct {
	EventID    string `gorm:"primaryKey"`
	Data       string
	WasRead    bool
	EventType  string
	OccurredOn time.Time
}

func (TransactionLogTrailing) TableName() string {
	return "transaction_log_trailing"
}
