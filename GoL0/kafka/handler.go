package kafka

import (
	dbase "GoL0/db"
	"database/sql"
	"encoding/json"
	"log"
)

func HandleMessage(db *sql.DB, data []byte) interface{} {
	var msg KafkaMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Println("Ошибка JSON:", err)
		return KafkaResponse{Error: "невалидный формат сообщения"}
	}

	switch msg.Type {
	case "query":
		var payload struct {
			OrderUID string `json:"order_uid"`
		}
		if err := json.Unmarshal(msg.Data, &payload); err != nil {
			log.Println("Невалидный запрос: ", err)
			return KafkaResponse{Error: "невалидный формат запроса"}
		}

		order, err := dbase.GetOrderById(db, payload.OrderUID)
		if err != nil {
			log.Println("Ошибка из БД:", err)
			return KafkaResponse{Error: err.Error()}
		}

		return KafkaResponse{Order: order}

	default:
		log.Println("Неизвестный тип:", msg.Type)
		return KafkaResponse{Error: "неподдерживаемый тип сообщения"}
	}
}
