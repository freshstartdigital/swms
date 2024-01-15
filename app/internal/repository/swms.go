package repository

import (
	"encoding/json"
	"log"

	"example.com/internal/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func (db *DB) GetAllSwms(OrganisationID int) ([]models.Swms, error) {

	rows, err := db.Query("SELECT id, user_id, name, swms_type, generator_status, file_name, file_path, created_at, updated_at, swms_data FROM swms WHERE organisation_id = $1", OrganisationID)

	if err != nil {
		log.Printf("Error getting all swms: %v", err)
		return nil, err
	}

	swms := []models.Swms{}
	for rows.Next() {
		var s models.Swms
		err := rows.Scan(&s.ID, &s.UserId, &s.Name, &s.SwmsType, &s.GeneratorStatus, &s.FileName, &s.FilePath, &s.CreatedAt, &s.UpdatedAt, &s.SwmsData)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		swms = append(swms, s)
	}

	defer rows.Close()
	return swms, nil
}

func (db *DB) CreateSwms(user models.Users, name string, SwmsData json.RawMessage, SwmsType string) (int, error) {

	var swmsID int

	stmt, err := db.Prepare("INSERT INTO swms (user_id, organisation_id, name, swms_type, swms_data) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		user.ID,
		user.OrganisationID,
		name,
		SwmsType,
		SwmsData,
	).Scan(&swmsID)

	if err != nil {
		return 0, err
	}

	return swmsID, nil
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
