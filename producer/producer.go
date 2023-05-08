package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// connect to kafka broker
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic1", 0)
	// set timer to stop trying to send message after 10 seconds
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	// write the message to be put in the specified topic
	conn.WriteMessages(kafka.Message{Value: []byte("Hello World")})
}
