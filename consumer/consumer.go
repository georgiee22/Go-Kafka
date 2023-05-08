package main

import (
	"context"
	"fmt"
	"time"

	kg "github.com/segmentio/kafka-go"
)

func main() {
	// connect to kafka broker
	conn, _ := kg.DialLeader(context.Background(), "tcp", "localhost:9092", "topic1", 0)
	// set timer to stop trying to send message after 8 seconds
	conn.SetReadDeadline(time.Now().Add(time.Second * 8))

	// message, _ := conn.ReadMessage(1e6)
	// fmt.Println(string(message.Value))

	batch_message := conn.ReadBatch(1e3, 1e6)
	bytes := make([]byte, 1e3)
	for {
		_, err := batch_message.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
}
