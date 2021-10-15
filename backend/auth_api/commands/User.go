package commands

import (
	"github.com/gofrs/uuid"
	"test1/core/core_events"
	"test1/core/event_store"
	"test1/core/message_queue"
)

func CreateUser(id uuid.UUID, username string, contactEmail string){
	userCreated := core_events.UserCreated{
		AggregateId: id,
		Username: username,
		ContactEmail: contactEmail,
	}
	message_queue.Publish(userCreated)
	event_store.PersistEvent(userCreated)
}
