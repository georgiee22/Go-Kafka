package kafka

import "github.com/segmentio/kafka-go"

var (
	Reader *kafka.Reader
)

func InitializeNewReader(broker, topic, groupid string) {
	// Set up a Kafka reader to read messages from the topic.
	Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupid,
	})
}
