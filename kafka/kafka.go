package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// TODO: Need to figure how to properly set up the kafka client in my application to manage the idea I had for it
// How will new clients connect to the app and act as a consumer, what will my func need to look like to keep the process open?
// What sort of business logic will need to be in the application to utilize the kafka client in this mannor. Or should I
// Just bake this logic into the client for now and decouple later.

// Not used atm
type Config struct {
	*kafka.ConfigMap
}

// Client represents a Kafka client.
type Client struct {
	producer *kafka.Producer
	consumer *kafka.Consumer
	// TODO: Add support for multiple 'consumers'
	// consumers map[string]*kafka.Consumer
}

type TopicConfig struct {
}

// NewClient creates a new Kafka client.
// TODO: Add depedency injection support.
func NewClient() (*Client, error) {
	newProducer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return nil, err
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{"bootstrap.servers": "localhost", "group.id": "group1"})

	if err != nil {
		log.Fatalf("error creating kafka client consumer: %v", err)
	}
	// defer consumer.Close()

	// TODO: Add support for multiple consumers
	// kafkaConsumers := make(map[string]*kafka.Consumer)

	return &Client{producer: newProducer, consumer: consumer}, nil
}

// Close closes the Kafka client.
func (c *Client) Close() {
	c.producer.Close()
	c.consumer.Close()
}

// ProduceMessage produces a message to the specified topic.
func (c *Client) ProduceMessage(topic string, message []byte) error {
	return c.producer.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic}, Value: message}, nil)
}

// ConsumeMessages consumes messages from the specified topic.
func (c *Client) ConsumeMessages(ctx context.Context, topic string, fromBeginning bool, handler func([]byte)) {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "group1",
	}
	consumer, err := kafka.NewConsumer(cfg)

	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	err = consumer.Subscribe(topic, nil)
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
			msg, err := consumer.ReadMessage(-1)
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}
			handler(msg.Value)
		}
	}
}

func (c *Client) NewTopic(ctx context.Context, name string, cfg TopicConfig) error {

	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		err := c.PublishMessage(ctx, name, word)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) PublishMessage(ctx context.Context, topic string, m string) error {
	err := c.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(m),
	}, nil)

	if err != nil {
		return err
	}

	return nil
}

// func (c *Client) RemoveTopic(ctx context.Context) error
func (c *Client) NewSubscription(ctx context.Context, topic string, ct string) error {
	err := c.consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return fmt.Errorf("failed to subscribe to topic: %v", err)
	}

	return nil
}

func (c *Client) NewConsumer(ctx context.Context) error

// type Producer interface {
// 	Produce(*kafka.Message, chan kafka.Event) error
// 	Close() error
// }

// // Consumer defines the interface for a Kafka consumer.
// type Consumer interface {
// 	SubscribeTopics([]string, kafka.RebalanceCb) error
// 	ReadMessage(int) (*kafka.Message, error)
// 	Close() error
// }

// type Config struct {
// 	bootstrapServer string
// 	brokers         []string
// }

// type Client struct {
// 	topics        []uuid.UUID // Meant to represent the 'topcis' within the Kafka instance
// 	rideProducer  *kafka.Producer
// 	rideConsumers map[string]*kafka.Consumer
// }
