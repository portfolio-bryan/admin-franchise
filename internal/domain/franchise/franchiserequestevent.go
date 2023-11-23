package franchise

import "github.com/bperezgo/admin_franchise/shared/domain/event"

const FranchiseRequestReceivedType event.Type = "events.franchise.requestreceived"

type FranchiseRequestReceivedEvent struct {
	event.BaseEvent
	id       string
	url      string
	duration string
}

func NewFranchiseRequestReceivedEvent(id, url string) FranchiseRequestReceivedEvent {
	return FranchiseRequestReceivedEvent{
		id:  id,
		url: url,

		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e FranchiseRequestReceivedEvent) Type() event.Type {
	return FranchiseRequestReceivedType
}

func (e FranchiseRequestReceivedEvent) CourseID() string {
	return e.id
}
