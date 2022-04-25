package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	conf := kafka.ReaderConfig{ // kafka reader configue
		Brokers: []string{"localhost:9092"},
		Topic:   "TestTopic",
		// GroupID:  "g1",
		// MaxBytes: 10,
	}
	reader := kafka.NewReader(conf) // created the kafka reader

	for {
		m, err := reader.ReadMessage(context.Background()) //this will read the message & give you a message or error
		if err != nil {
			fmt.Println("some error occured", err)
			continue
		}
		fmt.Println("Message is", string(m.Value))
	}
}
