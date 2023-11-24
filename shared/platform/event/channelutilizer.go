package event

import (
	"context"
	"log"

	"github.com/bperezgo/admin_franchise/shared/domain/event"
)

type ChannelUtilizer struct {
	handler       event.Handler
	channelError  ChannelError
	logTrailingDB LogTrailingDB
}

func NewChannelUtilizer(handler event.Handler, channelError ChannelError, logTrailingDB LogTrailingDB) ChannelUtilizer {
	return ChannelUtilizer{
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
			if ce.Error != nil {
				c.channelError.Publish(ce.Error)
				continue
			}

			if err := c.handler.Handle(ctx, ce.Event); err != nil {
				c.channelError.Publish(ce.Error)
				continue
			}

			if err := c.logTrailingDB.FulfillEvent(ce.Event); err != nil {
				c.channelError.Publish(ce.Error)
				continue
			}
		}
	}()
}
