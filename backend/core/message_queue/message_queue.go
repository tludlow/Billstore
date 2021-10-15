package message_queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"test1/core/core_events"
	"time"
)

func Consume(topic string, group string, execute func(m kafka.Message)){
	fmt.Println(topic + " started")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  group,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  time.Second * 1,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		execute(m)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func Publish(event core_events.IEvent){
	topic := "events-" + event.GetAggregateName()
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:   topic,
		Balancer: &kafka.LeastBytes{},
		BatchTimeout: time.Duration(1),
	}
	dataBytes, _ := json.Marshal(event)
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(event.GetType()),
			Value: dataBytes,
		},
	)
	if err != nil {
		if err == kafka.UnknownTopicOrPartition{
			//kafka.DialLeader(context.Background(), "tcp", "host:port", topic, partition)
			c, _ := kafka.Dial("tcp", "localhost:9092")
			kt := kafka.TopicConfig{Topic: topic, NumPartitions: 1, ReplicationFactor: 1}
			e := c.CreateTopics(kt)
			fmt.Println(e)
		}
		log.Println("Created new topic as one didn't exist")
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(event.GetType()),
				Value: dataBytes,
			},
		)
		if err != nil {log.Println("Error after creating new topic");return}
		//this needs to change, very bad
		Publish(event)
	}
	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}