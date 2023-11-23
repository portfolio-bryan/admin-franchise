package franchise

import (
	"encoding/json"

	"github.com/bperezgo/admin_franchise/shared/domain/event"
)

const FranchiseRequestReceivedType event.Type = "events.franchise.requestreceived"

type FranchiseRequestReceivedEventData struct {
	EventId string `json:"eventId"`
	Url     string `json:"url"`
}

type FranchiseRequestReceivedEvent struct {
	event.BaseEvent
	id       string
	url      string
	duration string
	data     FranchiseRequestReceivedEventData
}

func NewFranchiseRequestReceivedEvent(id, url string) FranchiseRequestReceivedEvent {
	baseEvent := event.NewBaseEvent(id)

	return FranchiseRequestReceivedEvent{
		id:  baseEvent.EventID(),
		url: url,

		BaseEvent: baseEvent,
	}
}

func (e FranchiseRequestReceivedEvent) Type() event.Type {
	return FranchiseRequestReceivedType
}

func (e FranchiseRequestReceivedEvent) Data() []byte {
	data := FranchiseRequestReceivedEventData{
		EventId: e.id,
		Url:     e.url,
	}

	b, err := json.Marshal(&data)
	if err != nil {
		return []byte{}
	}

	return b
}
