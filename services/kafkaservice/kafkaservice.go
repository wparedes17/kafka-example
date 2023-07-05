package kafkaservice

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaService interface {
	Write(ctx context.Context, key, value string) error
	Read(ctx context.Context) (string, error)
}

type kafkaObject struct {
	address []string
	topic   string
	groupid string
	writer  *kafka.Writer
	reader  *kafka.Reader
}

func New(topic, groupid string, address ...string) KafkaService {
	return &kafkaObject{
		address: address,
		topic:   topic,
		groupid: groupid,
	}
}

func (k *kafkaObject) Write(ctx context.Context, key, value string) error {
	k.writer = &kafka.Writer{
		Addr:     kafka.TCP(k.address...),
		Topic:    k.topic,
		Balancer: &kafka.LeastBytes{},
	}

	err := k.writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		},
	)

	k.writer.Close()

	return err
}

func (k *kafkaObject) Read(ctx context.Context) (string, error) {
	k.reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  k.address,
		Topic:    k.topic,
		GroupID:  k.groupid,
		MaxBytes: 10e6, // 10MB
	})

	m, err := k.reader.ReadMessage(ctx)
	if err != nil {
		return "", err
	}

	k.reader.Close()

	return string(m.Key) + "-" + string(m.Value), nil
}
