package core_events

import (
	"encoding/json"
	"github.com/EventStore/EventStore-Client-Go/messages"
	"github.com/gofrs/uuid"
)

type RefreshTokenCreated struct{
	AggregateId uuid.UUID
	UserIdentityId uuid.UUID
	UserId uuid.UUID
}
func (c RefreshTokenCreated) ToProposedEvent() messages.ProposedEvent{
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
func (c RefreshTokenCreated) GetAggregateId() uuid.UUID{
	return c.AggregateId
}
func (RefreshTokenCreated) GetAggregateName() string{
	return "RefreshToken"
}
func (RefreshTokenCreated) GetType() string{
	return "RefreshTokenCreated"
}

type RefreshTokenDeleted struct{
	AggregateId uuid.UUID
}
func (c RefreshTokenDeleted) ToProposedEvent() messages.ProposedEvent{
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
func (c RefreshTokenDeleted) GetAggregateId() uuid.UUID{
	return c.AggregateId
}
func (RefreshTokenDeleted) GetAggregateName() string{
	return "RefreshToken"
}
func (RefreshTokenDeleted) GetType() string{
	return "RefreshTokenDeleted"
}