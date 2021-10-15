package core_events

import (
	"github.com/EventStore/EventStore-Client-Go/messages"
	"github.com/gofrs/uuid"
)

type Event struct{
	AggregateId uuid.UUID
}

type IEvent interface{
	ToProposedEvent() messages.ProposedEvent
	GetAggregateName() string
	GetAggregateId() uuid.UUID
	GetType() string
}
