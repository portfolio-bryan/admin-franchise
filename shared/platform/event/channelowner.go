package event

import (
	"context"

	"github.com/bperezgo/admin_franchise/shared/domain/event"
)

type ChannelOwner struct {
	// With this repository we can handle async processes with the channels
	logTrailingDB LogTrailingDB

	channelError ChannelError

	channelEvents map[event.Type]chan ChannelEvent
}

func NewChannelOwner(logTrailingDB LogTrailingDB, channelError ChannelError) ChannelOwner {
	return ChannelOwner{
		logTrailingDB: logTrailingDB,
		channelError:  channelError,
		channelEvents: make(map[event.Type]chan ChannelEvent),
	}
}

func (c ChannelOwner) ChannelEvent(evtType event.Type) <-chan ChannelEvent {
	_, ok := c.channelEvents[evtType]
	if !ok {
		c.channelEvents[evtType] = make(chan ChannelEvent)
	}
	return c.channelEvents[evtType]
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

		channelEvent := c.channelEvents[evt.Type()]
		channelEvent <- ce
	}

	return nil
}
