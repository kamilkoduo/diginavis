package kafka_playground

import (
	"github.com/segmentio/kafka-go"
	"os"
)

type ReaderConfig struct {
	Topic string
}
type WriterConfig struct {
	Topic string
}

// Some Kafka-related variables
var Address = os.Getenv("KAFKA_BROKER_HOST")
var Partition = 0

// Topics declaration
// noinspection ALL
const (
	TOPIC_HELLO = "topic-hello"
)

func NewReader(config ReaderConfig) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{Address},
		Topic:    config.Topic,
		MinBytes: 1,    // 1B
		MaxBytes: 10e6, // 10MB
	})
}

func NewWriter(config WriterConfig) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{Address},
		Topic:   config.Topic,
	})
}
