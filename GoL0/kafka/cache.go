package kafka

import (
	dbase "GoL0/db"
	"database/sql"
	"log"
)

func RestoreCache(db *sql.DB) error {
	rows, err := db.Query(`SELECT order_uid FROM orders`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var uid string
		if err := rows.Scan(&uid); err != nil {
			log.Println("Ошибка чтения UID:", err)
			continue
		}
		order, err := dbase.GetOrderById(db, uid)
		if err != nil {
			log.Println("Ошибка восстановления заказа:", err)
			continue
		}
		Cache[uid] = *order
		log.Println("Кэш восстановлен:", uid)
	}
	return nil
}
