package db
package db

import (
	"database/sql"
	_"github.com/lib/pq"
)

func ConncetionDb() *sql.DB {
	connection := "user=teste dbname=AluraLoja password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}