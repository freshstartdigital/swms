package repository

import (
	"database/sql"
	"encoding/json"
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
	rows, err := db.Query("SELECT s.id, s.user_id, s.name, s.swms_type, s.generator_status, s.file_name, s.file_path, s.created_at, s.updated_at, sd.id, sd.swms_id, sd.data, sd.version FROM swms s JOIN swms_data sd ON s.id = sd.swms_id ORDER BY s.created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	swms := []models.Swms{}
	for rows.Next() {
		var s models.Swms
		var sd models.SwmsData
		err := rows.Scan(
			&s.ID,
			&s.UserId,
			&s.Name,
			&s.SwmsType,
			&s.GeneratorStatus,
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
			log.Printf("Error scanning SWMS row: %v\n", err)
			continue // or return nil, err
		}
		s.SwmsData = append(s.SwmsData, sd)
		swms = append(swms, s)
	}
	return swms, nil
}

func (db *DB) CreateSwms(user models.User, name string, data json.RawMessage) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	swms := models.Swms{
		Name:     name,
		SwmsType: "construction",
		SwmsData: []models.SwmsData{
			{
				Data:    data,
				Version: 1,
			},
		},
	}

	stmt, err := tx.Prepare("INSERT INTO swms(name, user_id, swms_type, created_at, updated_at) VALUES($1, $2, 'construction', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		name,
		user.ID,
	).Scan(&swms.ID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	stmt, err = tx.Prepare("INSERT INTO swms_data(swms_id, data, version) VALUES($1, $2, 1) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	for _, swmsData := range swms.SwmsData {
		err = stmt.QueryRow(
			swms.ID,
			swmsData.Data,
		).Scan(&swmsData.ID)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	tx.Commit()
	return swms.ID, nil
}

func (db *DB) UpdateFile(id int, file_name string, file_path string) error {

	stmt, err := db.Prepare("UPDATE swms SET file_name = $1, file_path = $2, generator_status = 'complete' WHERE id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		file_name,
		file_path,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
