package main

import (
	dbase "GoL0/db"
	"GoL0/handlers"
	"GoL0/kafka"
	"log"
	"net/http"
)

func main() {
	database := dbase.InitDB()
	defer database.Close()
	go kafka.StartDBConsumer(database)
	go kafka.StartServerConsumer(kafka.Callback)
	http.HandleFunc("/order", handlers.HandleOrderRequest)
	http.Handle("/", http.FileServer(http.Dir("./web")))

	log.Println("Сервис запущен на http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
