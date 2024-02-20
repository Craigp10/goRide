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

	t.Run("Create topic", func(t *testing.T) {
		// Create topic
		err := kc.NewTopic(ctx, "kafka.test.topic", TopicConfig{})

		require.NoError(t, err)
	})

	go kc.ConsumeMessages(ctx, "kafka.test.topic", true, func(m []byte) {})

	t.Run("Publish message", func(t *testing.T) {
		// Publish new message
		err := kc.PublishMessage(ctx, "kafka.test.topic", "craig is the best programmer")

		require.NoError(t, err)
	})

	// Need to wait for the messages to be published asynchronously
	time.Sleep(10 * time.Second)
	t.Run("Publish message", func(t *testing.T) {
		// Publish new message
		err := kc.PublishMessage(ctx, "kafka.test.topic", "craig is the best programmer")

		require.NoError(t, err)
	})
}
