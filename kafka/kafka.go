package kafka

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
)

// TODO: Need to figure how to properly set up the kafka client in my application to manage the idea I had for it
// How will new clients connect to the app and act as a consumer, what will my func need to look like to keep the process open?
// What sort of business logic will need to be in the application to utilize the kafka client in this mannor. Or should I
// Just bake this logic into the client for now and decouple later.

// Client represents a Kafka client.
type Client struct {
	producer  *kafka.Producer
	consumers map[string]*kafka.Consumer
}

type TopicConfig struct {
}

// NewClient creates a new Kafka client.
func NewClient(producerConfig *kafka.ConfigMap, consumerConfig *kafka.ConfigMap) (*Client, error) {
	newProducer, err := kafka.NewProducer(producerConfig)
	if err != nil {
		return nil, err
	}

	kafkaConsumers := make(map[string]*kafka.Consumer)

	return &Client{producer: newProducer, consumers: kafkaConsumers}, nil
}

// ProduceMessage produces a message to the specified topic.
func (c *Client) ProduceMessage(topic string, message []byte) error {
	return c.producer.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic}, Value: message}, nil)
}

// ConsumeMessages consumes messages from the specified topic.
func (c *Client) ConsumeMessages(ctx context.Context, topic string, fromBeginning bool, handler func([]byte)) {
	// c.NewSubscription(ctx, []string{topic}, nil)

	// for {
	// 	msg, err := c.consumers.ReadMessage(-1)
	// 	if err != nil {
	// 		// Handle error
	// 		continue
	// 	}

	// 	handler(msg.Value)
	// }
}

// Close closes the Kafka client.
func (c *Client) Close() {
	// c.consumers.Close()
	c.producer.Close()
}

func (c *Client) NewTopic(ctx context.Context, name string, cfg TopicConfig) error
func (c *Client) RemoveTopic(ctx context.Context) error
func (c *Client) NewSubscription(ctx context.Context, topicID uuid.UUID, ct string) error
func (c *Client) NewConsumer(ctx context.Context) error
