package event

import (
	"context"

	"github.com/bperezgo/admin_franchise/shared/domain/event"
)

type ChannelUtilizer struct {
	channelEvent <-chan ChannelEvent
	handler      event.Handler
}

func NewChannelUtilizer(handler event.Handler, channelEvent <-chan ChannelEvent) ChannelUtilizer {
	return ChannelUtilizer{
		handler:      handler,
		channelEvent: channelEvent,
	}
}

func (c ChannelUtilizer) Use(channelEvents <-chan ChannelEvent) {
	ctx := context.Background()
	select {
	case ce := <-c.channelEvent:
		c.handler.Handle(ctx, ce.Event)
	}
}
