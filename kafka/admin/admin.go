package admin

import (
	"context"
	"fmt"

	kafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var (
	defaultTimeout = 5000 // Milliseconds
)

// AdminClient defines the interface for managing Kafka topics and other administrative tasks.
type AdminClient interface {
	// CreateTopic creates a new topic in Kafka with the specified configuration.
	CreateTopic(config *TopicConfig) error

	// DeleteTopic deletes an existing topic in Kafka.
	DeleteTopic(topic string) error

	// DescribeTopic returns information about the specified topic.
	DescribeTopic(topic string) (*TopicConfig, error)

	// ListTopics returns a list of all topics in Kafka.
	ListTopics() ([]string, error)

	//...
}

type AdminClientConfig struct {
}

type KafkaAdminClient struct {
	ac *kafka.AdminClient
}

func newConfigMap(cfgMap AdminClientConfig) *kafka.ConfigMap {
	return &kafka.ConfigMap{"bootstrap.servers": "localhost"}
}
func NewAdminClient(configMap AdminClientConfig) (*KafkaAdminClient, error) {
	cfg := newConfigMap(configMap)
	adminClient, err := kafka.NewAdminClient(cfg)
	if err != nil {
		return nil, err
	}

	return &KafkaAdminClient{
		ac: adminClient,
	}, nil
}

// TopicConfig represents the configuration for a Kafka topic.
type TopicConfig struct {
	Topic             string
	NumPartitions     int
	ReplicationFactor int
	//...
}

func (kac *KafkaAdminClient) Close() error {
	kac.ac.Close()
	return nil
}

func (kac *KafkaAdminClient) CreateTopic(ctx context.Context, tc TopicConfig) (*kafka.TopicResult, error) {
	topicSpec := kafka.TopicSpecification{
		Topic:             tc.Topic,
		NumPartitions:     tc.NumPartitions,
		ReplicationFactor: tc.ReplicationFactor,
	}
	topicsSpec := []kafka.TopicSpecification{topicSpec}
	topics, err := kac.ac.CreateTopics(ctx, topicsSpec, nil)
	if err != nil {
		return nil, err
	}
	if len(topics) == 0 {
		return nil, fmt.Errorf("failed to create topic %s", tc.Topic)
	}

	return &topics[0], nil
}

type TopicPartition struct {
	parts []kafka.TopicPartitionInfo
}

type TopicDescribe struct {
	Topic                string
	Partitions           TopicPartition
	AuthorizedOperations []kafka.ACLOperation
}

// DescribeTopic describe a single topic within the kafka cluster
func (kac *KafkaAdminClient) DescribeTopic(ctx context.Context, topic string) (*TopicDescribe, error) {
	tc := kafka.NewTopicCollectionOfTopicNames([]string{topic})
	topics, err := kac.ac.DescribeTopics(ctx, tc, kafka.AdminOptionRequestTimeout{})
	if err != nil {
		return nil, err
	}

	topicsDescription := topics.TopicDescriptions

	if len(topicsDescription) < 1 {
		return nil, fmt.Errorf("topic not found to describe")
	}

	config := &TopicDescribe{
		Topic:                topicsDescription[0].Name,
		Partitions:           TopicPartition{parts: topicsDescription[0].Partitions},
		AuthorizedOperations: topicsDescription[0].AuthorizedOperations,
	}

	return config, nil
}

// DeleteTopic deletes an existing topic in Kafka.
func (kac *KafkaAdminClient) DeleteTopic(topic string) error {
	topics := []string{topic}
	_, err := kac.ac.DeleteTopics(context.Background(), topics)
	if err != nil {
		return err
	}

	return nil
}

// ListTopics returns a list of all topics in Kafka.
func (kac *KafkaAdminClient) ListTopics() ([]string, error) {
	topics, err := kac.ac.GetMetadata(nil, true, defaultTimeout)
	if err != nil {
		return nil, err
	}

	topicNames := make([]string, len(topics.Topics))
	j := 0
	for _, topic := range topics.Topics {
		topicNames[j] = topic.Topic
	}

	return topicNames, nil
}
