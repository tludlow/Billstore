package core_events

import (
	"encoding/json"
	"github.com/EventStore/EventStore-Client-Go/messages"
	"github.com/gofrs/uuid"
)

type EmailUserIdentityCreated struct{
	AggregateId uuid.UUID
	Email string
	UserId uuid.UUID
}
func (c EmailUserIdentityCreated) ToProposedEvent() messages.ProposedEvent{
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
func (c EmailUserIdentityCreated) GetAggregateId() uuid.UUID{
	return c.AggregateId
}
func (EmailUserIdentityCreated) GetAggregateName() string{
	return "EmailUserIdentity"
}
func (EmailUserIdentityCreated) GetType() string{
	return "EmailUserIdentityCreated"
}
//if we need default values, use constructor
//func NewEmailUserIdentityCreated(email string, userId uuid.UUID) EmailUserIdentityCreated{
//	return EmailUserIdentityCreated{
//		uuid.Must(uuid.NewV4()),
//		email,
//		userId,
//	}
//}