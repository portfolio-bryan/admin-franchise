package event

import "github.com/bperezgo/admin_franchise/shared/domain/event"

type ChannelEvent struct {
	Event event.Event
	Error error
}
