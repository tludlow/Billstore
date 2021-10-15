package commands

import (
	"github.com/gofrs/uuid"
	"test1/auth_api/hashing"
	"test1/auth_api/repo"
	"test1/core/core_events"
	"test1/core/domain"
	"test1/core/event_store"
	"test1/core/message_queue"
)

func CreateEmailUserIdentity(id uuid.UUID, email string, password string, userId uuid.UUID){
	userIdentity := domain.EmailUserIdentity{
	Id:       id,
	Email:    email,
	Password: hashing.HashPassword(password),
	UserId:   userId,
	}
	repo.AddEmailUserIdentity(userIdentity)
	identityCreated := core_events.EmailUserIdentityCreated{
	AggregateId: id,
	Email:       email,
	UserId:      userId,
	}
	message_queue.Publish(identityCreated)
	event_store.PersistEvent(identityCreated)
}
