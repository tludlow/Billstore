package core_events

import (
	"encoding/json"
	"github.com/EventStore/EventStore-Client-Go/messages"
	"github.com/gofrs/uuid"
)

type UserCreated struct {
	AggregateId uuid.UUID
	Username    string
	ContactEmail string
}
func (c UserCreated) ToProposedEvent() messages.ProposedEvent {
	dataBytes, _ := json.Marshal(c)
	event := messages.ProposedEvent{
		EventID:      uuid.Must(uuid.NewV4()),
		EventType:    c.GetType(),
		ContentType:  "application/json",
		Data:         dataBytes,
		UserMetadata: nil,
	}
	return event
}
func (c UserCreated) GetAggregateId() uuid.UUID{
	return c.AggregateId
}
func (UserCreated) GetAggregateName() string{
	return "User"
}
func (UserCreated) GetType() string {
	return "UserCreated"
}