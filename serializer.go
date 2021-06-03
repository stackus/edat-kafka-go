package edatkafkago

import (
	"github.com/segmentio/kafka-go"
	"github.com/stackus/edat/msg"
)

type Serializer interface {
	Serialize(message msg.Message) (kafka.Message, error)
	Deserialize(message kafka.Message) (msg.Message, error)
}

var DefaultSerializer = KafkaGoSerializer{}

type KafkaGoSerializer struct{}

var _ Serializer = (*KafkaGoSerializer)(nil)

func (s KafkaGoSerializer) Serialize(message msg.Message) (kafka.Message, error) {
	headers := make([]kafka.Header, 0, len(message.Headers()))

	for key, value := range message.Headers() {
		headers = append(headers, kafka.Header{
			Key:   key,
			Value: []byte(value),
		})
	}

	return kafka.Message{
		Value:   message.Payload(),
		Headers: headers,
	}, nil
}

func (s KafkaGoSerializer) Deserialize(message kafka.Message) (msg.Message, error) {
	var id string

	headers := make(msg.Headers, len(message.Headers))

	for _, header := range message.Headers {
		if header.Key == msg.MessageID {
			id = string(header.Value)
		} else {
			headers.Set(header.Key, string(header.Value))
		}
	}

	return msg.NewMessage(message.Value, msg.WithMessageID(id), msg.WithHeaders(headers)), nil
}
