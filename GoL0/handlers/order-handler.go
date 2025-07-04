package handlers

import (
	"GoL0/kafka"
	"encoding/json"
	"net/http"
	"time"
)

func HandleOrderRequest(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("id")
	if orderID == "" {
		http.Error(w, "Нет параметра 'id'", http.StatusBadRequest)
		return
	}

	if order, found := kafka.Cache[orderID]; found {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
		return
	}

	msg := kafka.KafkaMessage{
		Type: "query",
		Data: json.RawMessage(`{"order_uid":"` + orderID + `"}`),
	}
	err := kafka.Produce("order-in", msg)
	if err != nil {
		http.Error(w, "Kafka error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for i := 0; i < 20; i++ {
		if order, ok := kafka.Cache[orderID]; ok {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(order)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}

	http.Error(w, "Заказ не найден", http.StatusNotFound)
}
