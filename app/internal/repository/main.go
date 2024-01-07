package repository

import (
	"database/sql"
	"log"

	"example.com/internal/models"
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

func (db *DB) GetAllSwms() ([]models.Swms, error) {
	rows, err := db.Query("SELECT * FROM swms s JOIN swms_data sd ON s.id = sd.swms_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	swms := []models.Swms{}
	for rows.Next() {
		s := models.Swms{}
		sd := models.SwmsData{}
		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.SwmsType,
			&s.FileName,
			&s.FilePath,
			&s.CreatedAt,
			&s.UpdatedAt,
			&sd.ID,
			&sd.SwmsID,
			&sd.Data,
			&sd.Version,
		)
		if err != nil {
			return nil, err
		}
		s.SwmsData = append(s.SwmsData, sd)
		swms = append(swms, s)
	}
	return swms, nil
}
