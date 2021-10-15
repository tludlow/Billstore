package consumers

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"test1/core/core_events"
	"test1/core/message_queue"
	"test1/user_api/repo"
	"test1/user_api/view_models"
)

func Start(){
	go message_queue.Consume("events-User", "user-events", func(m kafka.Message){
		event := new(core_events.UserCreated)
		if err := json.Unmarshal(m.Value, event); err != nil{log.Println(err)}
		userDetails := view_models.UserDetails{
			Id:             event.AggregateId,
			Username: 		event.Username,
			ContactEmail: event.ContactEmail,
			AgeVerified: false,
			MatureContentFilter: false,
		}
		repo.AddUserDetails(userDetails)
	})
}