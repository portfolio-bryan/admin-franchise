package notify

import (
	"context"
	"log"

	"github.com/bperezgo/admin_franchise/shared/domain/event"
)

type NotificationHandler struct{}

func NewNotificationHandler() NotificationHandler {
	return NotificationHandler{}
}

func (h NotificationHandler) Handle(ctx context.Context, evt event.Event) error {
	log.Println("NotificationHandler")
	return nil
}
