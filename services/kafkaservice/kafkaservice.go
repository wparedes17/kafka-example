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
	writer *kafka.Writer
	reader *kafka.Reader
}

func New(topic string, address ...string) KafkaService {
	return &kafkaObject{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(address...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:   address,
			Topic:     topic,
			Partition: 0,
			MaxBytes:  10e6, // 10MB
		}),
	}
}

func (k *kafkaObject) Write(ctx context.Context, key, value string) error {
	err := k.writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		},
	)

	return err
}

func (k *kafkaObject) Read(ctx context.Context) (string, error) {
	m, err := k.reader.ReadMessage(ctx)
	if err != nil {
		return "", err
	}

	return string(m.Value), nil
}
