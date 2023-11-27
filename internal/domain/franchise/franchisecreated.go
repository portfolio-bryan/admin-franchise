package franchise

import (
	"encoding/json"

	"github.com/bperezgo/admin_franchise/shared/domain/event"
)

const FranchiseCreatedType event.Type = "events.franchise.created"

type FranchiseCreatedEvent struct {
	event.BaseEvent
	data      FranchiseDTO
	eventType event.Type
}

func NewFranchiseCreatedEvent(franchiseDTO FranchiseDTO) FranchiseCreatedEvent {
	return FranchiseCreatedEvent{
		data:      franchiseDTO,
		eventType: FranchiseCreatedType,

		BaseEvent: event.NewBaseEvent(franchiseDTO.ID),
	}
}

func (e FranchiseCreatedEvent) Type() event.Type {
	return e.eventType
}

func (e FranchiseCreatedEvent) Data() []byte {
	b, err := json.Marshal(&e.data)
	if err != nil {
		return []byte{}
	}

	return b
}
