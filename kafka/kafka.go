package kafka

import (
	"RouteRaydar/kafka/admin"
	consumer "RouteRaydar/kafka/consumer"
	producer "RouteRaydar/kafka/producer"
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// TODO: Need to figure how to properly set up the kafka client in the application to manage the idea I had for it
// How will new clients connect to the app and act as a consumer, what will the func need to look like to keep the process open?
// What sort of business logic will need to be in the application to utilize the kafka client in this mannor. should
// just bake this logic into the client for now and decouple later?

// Not used atm
type Config struct {
	*kafka.ConfigMap
}

// Client represents a Kafka client.
type Client struct {
	producer    *producer.KafkaProducer
	consumer    *consumer.KafkaConsumer
	adminClient *admin.KafkaAdminClient
	// TODO: Add support for multiple 'consumers'
	// consumers map[string]*kafka.Consumer
}

type TopicConfig struct{}

type TopicDetails struct {
	admin.TopicDescribe
}

// NewClient creates a new Kafka client.
// TODO: Add depedency injection support.
func NewClient() (*Client, error) {
	newProducer, err := producer.NewProducer()
	if err != nil {
		return nil, err
	}

	consumer, err := consumer.NewConsumer()

	if err != nil {
		log.Fatalf("error creating kafka client consumer: %v", err)
	}
	// defer consumer.Close()

	// TODO: Add support for multiple consumers
	// kafkaConsumers := make(map[string]*kafka.Consumer)
	cfg := admin.AdminClientConfig{}

	ac, err := admin.NewAdminClient(cfg)

	if err != nil {
		log.Fatalf("error creating kafka client consumer: %v", err)
	}

	return &Client{producer: newProducer, consumer: consumer, adminClient: ac}, nil
}

// Close closes the Kafka client.
func (c *Client) Close() {
	c.producer.Close()
	c.consumer.Close()
}

// ProduceMessage produces a message to the specified topic.
func (c *Client) ProduceMessage(topic string, message []byte) error {
	return c.producer.ProduceMessage(topic, message)
}

func (c *Client) NewTopic(ctx context.Context, name string, cfg TopicConfig) error {
	tc := admin.TopicConfig{
		Topic:             "kafka.test.topic",
		NumPartitions:     1,
		ReplicationFactor: 1,
	}
	_, err := c.adminClient.CreateTopic(ctx, tc)
	if err != nil {
		return err
	}

	return nil
}

// func (c *Client) DescribeTopic(ctx context.Context, name string) (kafka.TopicDescription, error) {
// 	var topic kafka.TopicDescription

// 	return topic, nil
// }

func (c *Client) PublishMessage(ctx context.Context, topic string, m string) error {
	err := c.producer.ProduceMessage(topic, []byte(m))

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

func (c *Client) DescribeTopic(ctx context.Context, topic string) (*TopicDetails, error) {
	to, err := c.adminClient.DescribeTopic(ctx, topic)
	if err != nil {
		return nil, err
	}

	return &TopicDetails{
		*to,
	}, nil
}

// type Config struct {
// 	bootstrapServer string
// 	brokers         []string
// }
