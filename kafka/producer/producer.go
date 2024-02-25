package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// The Producer file intends to develop a producer interface that can be interchangable of kafka clients.
// Aimed to handle any producer specific logic, like partitions.
type Producer interface {
	ProduceMessage(topic string, m []byte, ch chan kafka.Event) error
	Close() error
}

type KafkaProducer struct {
	// Producer
	producer *kafka.Producer
}

type ProducerConfig struct {
	// Producer
	broker string
}

func NewProducerConfig(broker string) *ProducerConfig {
	return &ProducerConfig{
		broker: broker,
	}
}
func NewProducer(producerConfig *ProducerConfig) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": producerConfig.broker})
	// p, err := kafka.NewProducer(producerConfig.Config)
	if err != nil {
		return nil, fmt.Errorf("error creating producer: %v", err)
	}

	kp := &KafkaProducer{
		producer: p,
	}

	return kp, nil
}

func (p *KafkaProducer) ProduceMessage(topic string, m []byte) error {
	km := &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic}, Value: m}
	err := p.producer.Produce(km, nil)
	if err != nil {
		return fmt.Errorf("error producing message: %v", err)
	}

	return err
}

func (p *KafkaProducer) Close() error {
	p.producer.Close()
	return nil
}
