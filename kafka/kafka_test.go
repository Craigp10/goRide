package kafka

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
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
		// Length is 1 - consumer_offset -- kafka.test.topic.1, kafka.test.topic.2 have no messages so aren't considered yet?
		require.Equal(t, 1, len(topics))
		t.Log(topics)
	})

	var id *uuid.UUID
	t.Run("Subscribe to topic", func(t *testing.T) {
		id, err = kc.NewSubscription(ctx, topic)
		require.NoError(t, err)
		require.NotEmpty(t, id)
	})

	t.Run("Confirm consumer exist in map", func(t *testing.T) {
		count := kc.GetConsumerCount()
		require.Equal(t, 1, count)
	})

	go kc.consumersMap[*id].ConsumeMessages(ctx, topic, true, func(m []byte) {})

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
