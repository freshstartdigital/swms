package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type DB struct {
	*sql.DB
}

func ConnectDB() *DB {
	db, err := sql.Open("postgres", "host=db port=5432 user=your_user password=your_password dbname=your_db_name sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return &DB{db}
}

func (db *DB) Close() error {
	return db.DB.Close()
}
