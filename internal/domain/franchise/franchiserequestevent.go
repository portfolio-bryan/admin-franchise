package franchise

import (
	"encoding/json"

	"github.com/bperezgo/admin_franchise/shared/domain/event"
)

const FranchiseRequestReceivedType event.Type = "events.franchise.requestreceived"

type FranchiseRequestReceivedEventData struct {
	AggregateID string `json:"aggregateId"`
	EventId     string `json:"eventId"`
	Url         string `json:"url"`
}

type FranchiseRequestReceivedEvent struct {
	event.BaseEvent
	id       string
	eventID  string
	url      string
	duration string
	data     FranchiseRequestReceivedEventData
}

func NewFranchiseRequestReceivedEvent(aggregateID, url string) FranchiseRequestReceivedEvent {
	baseEvent := event.NewBaseEvent(aggregateID)

	return FranchiseRequestReceivedEvent{
		id:      aggregateID,
		url:     url,
		eventID: baseEvent.EventID(),

		BaseEvent: baseEvent,
	}
}

func (e FranchiseRequestReceivedEvent) Type() event.Type {
	return FranchiseRequestReceivedType
}

func (e FranchiseRequestReceivedEvent) Data() []byte {
	data := FranchiseRequestReceivedEventData{
		AggregateID: e.id,
		EventId:     e.eventID,
		Url:         e.url,
	}

	b, err := json.Marshal(&data)
	if err != nil {
		return []byte{}
	}

	return b
}
