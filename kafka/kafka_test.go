package kafka

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestKafkaInit(t *testing.T) {
	ctx := context.Background()

	var kc *Client
	var err error
	t.Run("Initialize Kafka Client", func(t *testing.T) {
		kc, err = NewClient()
		require.NoError(t, err)
		require.NotEmpty(t, kc)
	})
	defer kc.Close()
	topic := "kafka.test.topic.1"
	t.Run("Create topic", func(t *testing.T) {
		// Create topic
		err := kc.NewTopic(ctx, topic, TopicConfig{})

		require.NoError(t, err)
	})
	// t.Run("New topic", func(t *testing.T) {
	// 	// Create topic
	// 	err := kc.NewTopic(ctx, "kafka.test.topic.2", TopicConfig{})

	// 	require.NoError(t, err)
	// })

	t.Run("Describe Topic", func(t *testing.T) {
		topic, err := kc.DescribeTopic(ctx, "kafka.test.topic.2")
		require.NoError(t, err)
		require.NotEmpty(t, topic)
	})

	time.Sleep(10 * time.Second)
	t.Run("List Topics", func(t *testing.T) {
		topics, err := kc.adminClient.ListTopics()
		require.NoError(t, err)
		require.NotEmpty(t, topics)
		// Length is 3 - consumer_offset, kafka.test.topic.1, kafka.test.topic.2
		require.Equal(t, 2, len(topics))
	})

	go kc.consumer.ConsumeMessages(ctx, topic, true, func(m []byte) {})

	t.Run("Publish message", func(t *testing.T) {
		// Publish new message
		err := kc.PublishMessage(ctx, topic, "craig is the best programmer")
		require.NoError(t, err)
	})

	// Need to wait for the messages to be published asynchronously
	time.Sleep(10 * time.Second)
	t.Run("Publish message", func(t *testing.T) {
		// Publish new message
		err := kc.PublishMessage(ctx, topic, "Cam is the best programmer")
		require.NoError(t, err)
	})
}
