package event

import (
	"context"

	"github.com/bperezgo/admin_franchise/shared/domain/event"
)

type ChannelOwner struct {
	// TODO: Use an strategy to handle many channels and subscriptions
	channelEvents chan ChannelEvent
	// TODO: Use Some table of postgres to this
	logTrailingDB interface{}
}

func NewChannelOwner() ChannelOwner {
	channelEvents := make(chan ChannelEvent)
	return ChannelOwner{
		channelEvents: channelEvents,
	}
}

func (c ChannelOwner) ChannelEvents() <-chan ChannelEvent {
	return c.channelEvents
}

func (c ChannelOwner) Publish(ctx context.Context, events []event.Event) error {
	for _, evt := range events {
		// TODO: Write in the table of postgres that the event is received by the event bus

		ce := ChannelEvent{
			Event: evt,
		}

		c.channelEvents <- ce
	}

	return nil
}
