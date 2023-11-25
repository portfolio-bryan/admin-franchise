package event

import (
	"context"

	"github.com/bperezgo/admin_franchise/shared/domain/event"
	"github.com/google/uuid"
)

type ChannelID string

type ChannelOwner struct {
	// With this repository we can handle async processes with the channels
	logTrailingDB LogTrailingDB

	channelError ChannelError

	channelEvents map[event.Type]map[ChannelID]chan ChannelEvent
}

func NewChannelOwner(logTrailingDB LogTrailingDB, channelError ChannelError) ChannelOwner {

	return ChannelOwner{
		logTrailingDB: logTrailingDB,
		channelError:  channelError,
		channelEvents: make(map[event.Type]map[ChannelID]chan ChannelEvent),
	}
}

// Posible bug: if the channel utilizer is repeated in the injection, it will be repeated innecessary the message
func (c ChannelOwner) GetChannel(evtType event.Type, channelID ChannelID) <-chan ChannelEvent {
	_, ok := c.channelEvents[evtType]
	if !ok {
		c.channelEvents[evtType] = make(map[ChannelID]chan ChannelEvent)
	}
	_, ok = c.channelEvents[evtType][channelID]
	if !ok {
		c.channelEvents[evtType][channelID] = make(chan ChannelEvent)
	}
	return c.channelEvents[evtType][channelID]
}

func (c ChannelOwner) Subscribe(channelUtilizer ChannelUtilizer) ChannelID {
	return ChannelID(uuid.NewString())
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

		channelEvents := c.channelEvents[evt.Type()]
		for _, channelEvent := range channelEvents {
			channelEvent <- ce
		}
	}

	return nil
}
