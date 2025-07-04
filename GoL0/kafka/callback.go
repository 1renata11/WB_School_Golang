package kafka

import (
	"GoL0/models"
	"encoding/json"
	"log"
)

var Cache = make(map[string]models.Order)

func Callback(msg []byte) {
	var resp KafkaResponse
	if err := json.Unmarshal(msg, &resp); err != nil {
		log.Println("Ошибка при разборе ответа:", err)
		return
	}

	if resp.Error != "" {
		log.Println("Ошибка от сервиса:", resp.Error)
		return
	}

	if resp.Order != nil {
		Cache[resp.Order.OrderUID] = *resp.Order
		log.Printf("Ответ добавлен в кэш: %s", resp.Order.OrderUID)
	}
}
