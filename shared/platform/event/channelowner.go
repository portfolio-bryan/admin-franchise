package event

import (
	"context"

	"github.com/bperezgo/admin_franchise/shared/domain/event"
)

type ChannelOwner struct {
	// TODO: Use an strategy to handle many channels and subscriptions
	channelEvents chan ChannelEvent
	// With this repository we can handle async processes with the channels
	logTrailingDB LogTrailingDB

	channelError ChannelError
}

func NewChannelOwner(logTrailingDB LogTrailingDB, channelError ChannelError) ChannelOwner {
	channelEvents := make(chan ChannelEvent)
	return ChannelOwner{
		channelEvents: channelEvents,
		logTrailingDB: logTrailingDB,
	}
}

func (c ChannelOwner) ChannelEvents() <-chan ChannelEvent {
	return c.channelEvents
}

func (c ChannelOwner) Publish(ctx context.Context, events []event.Event) error {
	for _, evt := range events {
		if err := c.logTrailingDB.SavePendingEvent(evt); err != nil {
			c.channelError.Publish(err)
			continue
		}

		ce := ChannelEvent{
			Event: evt,
		}

		c.channelEvents <- ce
	}

	return nil
}
