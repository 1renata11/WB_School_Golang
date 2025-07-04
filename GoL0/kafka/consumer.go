package kafka

import (
	"context"
	"database/sql"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartDBConsumer(db *sql.DB) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "order-in",
		GroupID: "order-group-db",
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Ошибка Kafka:", err)
		}

		answer := HandleMessage(db, msg.Value)
		if answer != nil {
			Produce("order-out", answer)
		}
	}
}

func StartServerConsumer(callback func([]byte)) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "order-out",
		GroupID: "order-group-server",
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Ошибка Kafka:", err)
			continue
		}

		callback(msg.Value)
	}
}
