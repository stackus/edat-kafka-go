package edatkafkago

import (
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stackus/edat/log"
)

type ConsumerOption func(*Consumer)

func WithConsumerAckWait(ackWait time.Duration) ConsumerOption {
	return func(consumer *Consumer) {
		consumer.ackWait = ackWait
	}
}

func WithConsumerSerializer(serializer Serializer) ConsumerOption {
	return func(consumer *Consumer) {
		consumer.serializer = serializer
	}
}

func WithConsumerLogger(logger log.Logger) ConsumerOption {
	return func(consumer *Consumer) {
		consumer.logger = logger
	}
}

func WithConsumerDialer(dialer *kafka.Dialer) ConsumerOption {
	return func(consumer *Consumer) {
		consumer.dialer = dialer
	}
}
