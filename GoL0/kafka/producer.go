package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

func Produce(topic string, value interface{}) error {
	writer := kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: topic,
	}
	data, _ := json.Marshal(value)
	return writer.WriteMessages(context.Background(), kafka.Message{
		Value: data,
	})
}
