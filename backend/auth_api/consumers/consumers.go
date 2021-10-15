package consumers

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"test1/auth_api/repo"
	"test1/core/core_events"
	"test1/core/domain"
	"test1/core/message_queue"
)

func Start(){
	go message_queue.Consume("events-EmailUserIdentity", "emailUserIdentity-events", func(m kafka.Message){
		fmt.Println("h2")
	})
	go message_queue.Consume("events-RefreshToken", "refreshToken-events", func(m kafka.Message){
		//add switch for if created or deleted
		switch string(m.Key) {
		case "RefreshTokenCreated":
			event := new(core_events.RefreshTokenCreated)
			if err := json.Unmarshal(m.Value, event); err != nil{
				log.Println(err)
			}
			refreshToken := domain.RefreshToken{
				Id:             event.AggregateId,
				UserId: 		event.UserId,
				UserIdentityId: event.UserIdentityId,
			}
			repo.AddRefreshToken(refreshToken)
			fmt.Println("h3")
		case "RefreshTokenDeleted":
			event := new(core_events.RefreshTokenCreated)
			if err := json.Unmarshal(m.Value, event); err != nil{
				log.Println(err)
			}
			repo.DeleteRefreshTokenById(event.AggregateId)
		}
	})
}