package kafka

import (
	"GoL0/models"
	"encoding/json"
)

type KafkaMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type KafkaResponse struct {
	Order *models.Order `json:"order,omitempty"`
	Error string        `json:"error,omitempty"`
}
