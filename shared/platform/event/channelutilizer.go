package event

import (
	"context"
	"log"

	"github.com/bperezgo/admin_franchise/shared/domain/event"
	"github.com/google/uuid"
)

type ChannelID string

type ChannelUtilizer struct {
	channelID     ChannelID
	handler       event.Handler
	channelError  ChannelError
	logTrailingDB LogTrailingDB
}

func NewChannelUtilizer(handler event.Handler, channelError ChannelError, logTrailingDB LogTrailingDB) ChannelUtilizer {
	return ChannelUtilizer{
		channelID:     ChannelID(uuid.NewString()),
		handler:       handler,
		channelError:  channelError,
		logTrailingDB: logTrailingDB,
	}
}

func (c ChannelUtilizer) Use(channelEvent <-chan ChannelEvent) {
	ctx := context.Background()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic occurred in channel utilizer, the go function is rerun again:", err)
			}
		}()

		for ce := range channelEvent {
			go func(ce ChannelEvent) {
				defer func() {
					if err := recover(); err != nil {
						log.Println("panic occurred in channel utilizer, the go function is rerun again:", err)
					}
				}()

				if ce.Error != nil {
					c.channelError.Publish(ce.Error)
					return
				}

				if err := c.handler.Handle(ctx, ce.Event); err != nil {
					c.channelError.Publish(ce.Error)
					return
				}

				if err := c.logTrailingDB.FulfillEvent(ce.Event); err != nil {
					c.channelError.Publish(ce.Error)
					return
				}
			}(ce)
		}
	}()
}
