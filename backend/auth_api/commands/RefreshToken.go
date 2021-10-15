package commands

import (
	"github.com/gofrs/uuid"
	"test1/auth_api/repo"
	"test1/core/core_events"
	"test1/core/event_store"
	"test1/core/message_queue"
)

func CreateRefreshToken(id uuid.UUID, userIdentityId uuid.UUID, userId uuid.UUID){
	refreshTokenCreated := core_events.RefreshTokenCreated{AggregateId: id, UserIdentityId: userIdentityId, UserId: userId}
	message_queue.Publish(refreshTokenCreated)
	event_store.PersistEvent(refreshTokenCreated)
}

func DeleteRefreshToken(id uuid.UUID){
	refreshTokenDeleted := core_events.RefreshTokenDeleted{AggregateId: id}
	message_queue.Publish(refreshTokenDeleted)
	event_store.PersistEvent(refreshTokenDeleted)
}

// DeleteAllRefreshTokens Don't really want to publish this, not going to be consumed by anyone
func DeleteAllRefreshTokens(userId uuid.UUID) int{
	tokens := repo.GetAllRefreshTokensForUser(userId)
	events := make([]core_events.IEvent, len(tokens))
	for i, token := range tokens {
		events[i] = core_events.RefreshTokenDeleted{AggregateId: token.Id}
	}
	count := int(repo.DeleteAllRefreshTokensForUser(userId))
	event_store.PersistEvents(events)
	return count
}