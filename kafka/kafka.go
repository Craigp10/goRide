package kafka

import (
	"RouteRaydar/kafka/admin"
	consumer "RouteRaydar/kafka/consumer"
	producer "RouteRaydar/kafka/producer"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

// TODO: Need to figure how to properly set up the kafka client in the application to manage the idea I had for it
// How will new clients connect to the app and act as a consumer, what will the func need to look like to keep the process open?
// What sort of business logic will need to be in the application to utilize the kafka client in this mannor. should
// just bake this logic into the client for now and decouple later?
const (
	BOOTSTRAP_SERVER = "localhost"
)

type Config struct {
	// consumerGroups string
	ProducerConfig *producer.ProducerConfig
	ConsumerConfig *consumer.ConsumerConfig
}

// Client represents a Kafka client.
type Client struct {
	producer     *producer.KafkaProducer
	adminClient  *admin.KafkaAdminClient
	consumersMap map[uuid.UUID]*consumer.KafkaConsumer
	consumers    int
}

type TopicConfig struct{}

type TopicDetails struct {
	admin.TopicDescribe
}

// NewClient creates a new Kafka client.
// TODO: Add depedency injection support.
func NewClient() (*Client, error) {
	cfg := producer.NewProducerConfig(BOOTSTRAP_SERVER)
	newProducer, err := producer.NewProducer(cfg)
	if err != nil {
		return nil, err
	}

	// consumer, err := consumer.NewConsumer()
	consumers := 0

	consumersMap := make(map[uuid.UUID]*consumer.KafkaConsumer)

	if err != nil {
		log.Fatalf("error creating kafka client consumer: %v", err)
	}
	// defer consumer.Close()

	adminCfg := admin.AdminClientConfig{}

	ac, err := admin.NewAdminClient(adminCfg)

	if err != nil {
		log.Fatalf("error creating kafka client consumer: %v", err)
	}

	return &Client{producer: newProducer, consumersMap: consumersMap, consumers: consumers, adminClient: ac}, nil
}

// Close closes the Kafka client. Panics on error, will need to update later.
func (c *Client) Close() {
	err := c.producer.Close()
	if err != nil {
		log.Panicf("error closing kafka producer: %v", err)
	}
	for consumer := range c.consumersMap {
		err = c.consumersMap[consumer].Close()
		if err != nil {
			log.Panicf("error closing kafka consumer: %s with error: %v", consumer, err)
		}
		c.consumers--
	}
}

func (c *Client) GetConsumerCount() int {
	return c.consumers
}

// ProduceMessage produces a message to the specified topic.
func (c *Client) ProduceMessage(topic string, message []byte) error {
	return c.producer.ProduceMessage(topic, message)
}

func (c *Client) NewTopic(ctx context.Context, name string) error {
	tc := admin.TopicConfig{
		Topic:             name,
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
func (c *Client) NewSubscription(ctx context.Context, topic string, cfg *consumer.ConsumerConfig) (*uuid.UUID, error) {
	newConsumer, err := consumer.NewConsumer(cfg)
	if err != nil {
		return nil, fmt.Errorf("error generating new consumer in new subscription")
	}
	err = newConsumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {

		return nil, fmt.Errorf("failed to subscribe to topic: %v", err)
	}
	newConsumerId := uuid.New()
	c.consumersMap[newConsumerId] = newConsumer
	c.consumers++

	return &newConsumerId, nil
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
