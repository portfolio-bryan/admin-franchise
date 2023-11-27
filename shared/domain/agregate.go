package domain

import "github.com/bperezgo/admin_franchise/shared/domain/event"

type Aggregate struct {
	Data      interface{}
	EventType event.Type

	events []event.Event
}

func NewAggregate(data interface{}, eventType event.Type) Aggregate {
	return Aggregate{
		Data:      data,
		EventType: eventType,
	}
}

// Record records a new domain event.
func (a Aggregate) Record(evt event.Event) {
	a.events = append(a.events, evt)
}

// PullEvents returns all the recorded domain events.
func (a Aggregate) PullEvents() []event.Event {
	evt := a.events
	a.events = []event.Event{}

	return evt
}
