package event_store

import (
	"context"
	"fmt"
	"github.com/EventStore/EventStore-Client-Go/client"
	"github.com/EventStore/EventStore-Client-Go/messages"
	"github.com/EventStore/EventStore-Client-Go/streamrevision"
	"test1/core/core_events"
	"time"
)

func PersistEvents(eventList []core_events.IEvent) error{
	connectionString, err := client.ParseConnectionString("esdb://127.0.0.1:2113?tls=false")
	if err != nil {return err}
	eventStoreClient, err := client.NewClient(connectionString)
	if err != nil {return err}
	defer eventStoreClient.Close()
	connection, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	eventByAggregate := make(map[string][]messages.ProposedEvent)
	for _, event := range eventList {
		streamId := fmt.Sprintf("%s-%s", event.GetAggregateName(), event.GetAggregateId().String())
		eventByAggregate[streamId] = append(eventByAggregate[streamId], event.ToProposedEvent())
	}
	for streamId, proposedEvents := range eventByAggregate {
		_, err = eventStoreClient.AppendToStream(connection, streamId, streamrevision.StreamRevisionAny, proposedEvents)
		if err != nil {return err}
	}
	return nil

}

func PersistEvent(event core_events.IEvent) error{
	connectionString, err := client.ParseConnectionString("esdb://127.0.0.1:2113?tls=false")
	if err != nil {return err}
	eventStoreClient, err := client.NewClient(connectionString)
	if err != nil {return err}
	defer eventStoreClient.Close()
	connection, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	streamId := fmt.Sprintf("%s-%s", event.GetAggregateName(), event.GetAggregateId().String())
	proposedEvent := []messages.ProposedEvent{event.ToProposedEvent()}
	_, err = eventStoreClient.AppendToStream(connection, streamId, streamrevision.StreamRevisionAny, proposedEvent)
	if err != nil {return err}
	return nil
}

func PersistEventsSingleAggregate(eventList []core_events.IEvent) error{
	connectionString, err := client.ParseConnectionString("esdb://127.0.0.1:2113?tls=false")
	if err != nil {return err}
	eventStoreClient, err := client.NewClient(connectionString)
	if err != nil {return err}
	defer eventStoreClient.Close()
	connection, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	proposedEvents := make([]messages.ProposedEvent, len(eventList))
	for i, event := range eventList {
		proposedEvents[i] = event.ToProposedEvent()
	}
	streamId := fmt.Sprintf("%s-%s", eventList[0].GetAggregateName(), eventList[0].GetAggregateId().String())
	_, err = eventStoreClient.AppendToStream(connection, streamId, streamrevision.StreamRevisionAny, proposedEvents)
	return nil
}