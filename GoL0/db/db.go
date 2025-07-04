package dbase

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := "postgres://user:password@localhost:5433/mydb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка открытия mydb: %v\n", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка подключения к mydb: %v", err)
	}
	log.Println("Успешное подключение к mydb")
	return db
}
