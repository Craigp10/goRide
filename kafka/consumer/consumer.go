package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Consumer defines the interface for a Kafka consumer.
type Consumer interface {
	SubscribeTopics([]string, kafka.RebalanceCb) error
	ReadMessage(int) (*kafka.Message, error)
	Close() error
	// *kafka.Consumer
}

type KafkaConsumer struct {
	Consumer
	consumer *kafka.Consumer
}

func NewConsumer() (*KafkaConsumer, error) {
	kc, err := kafka.NewConsumer(&kafka.ConfigMap{"bootstrap.servers": "localhost", "group.id": "group1"})
	if err != nil {
		return nil, fmt.Errorf("error creating consumer: %v", err)
	}

	c := &KafkaConsumer{
		consumer: kc,
	}

	return c, nil
}

func (c *KafkaConsumer) Close() error {
	c.consumer.Close()
	return nil
}

func (c *KafkaConsumer) SubscribeTopics(topics []string, rb kafka.RebalanceCb) error {

	return nil
}

// ConsumeMessages consumes messages from the specified topic.
func (kc *KafkaConsumer) ConsumeMessages(ctx context.Context, topic string, fromBeginning bool, handler func([]byte)) {
	// TODO: Look into Consumers Config mapping
	// cfg := &kafka.ConfigMap{
	// 	"bootstrap.servers": "localhost",
	// 	"group.id":          "group1",
	// }
	err := kc.consumer.Subscribe(topic, nil)

	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	// Assumes consumer wants to close after consuming messages
	// defer consumer.Close()

	// err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topic %s: %v", topic, err)
	}

	fmt.Println("Consumer is listening")
	for {
		select {
		case <-ctx.Done():
			log.Println("Context cancelled. Stopping consumer.")
			return
		default:
			msg, err := kc.consumer.ReadMessage(-1)
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}
			handler(msg.Value)
		}
	}
}
