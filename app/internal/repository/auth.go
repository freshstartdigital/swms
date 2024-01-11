package repository

import (
	"log"

	"example.com/internal/models"
)

func (db *DB) GetUser(email string, password string) (models.Users, error) {
	var user models.Users
	err := db.QueryRow(`
		SELECT id, organisation_id, name, email, password, created_at, updated_at 
		FROM users 
		WHERE email = $1 AND password = $2`, email, password).
		Scan(&user.ID, &user.OrganisationID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}

func (db *DB) GetOrganisation(id int) (models.Organisations, error) {
	var organisation models.Organisations
	var organisationType models.OrganisationTypes
	err := db.QueryRow(`
		SELECT
			org.id, org.name, org.business_address, org.abn, org.business_phone, 
			org.business_email, org.logo_file_name, org.created_at, org.updated_at, 
			org.account_holder_id, org_type.id, org_type.type, org_type.display_name 
		FROM organisations org
		JOIN organisation_types org_type ON org.organisation_type_id = org_type.id
		WHERE org.id = $1`, id).
		Scan(&organisation.ID, &organisation.Name, &organisation.BusinessAddress, &organisation.ABN,
			&organisation.BusinessPhone, &organisation.BusinessEmail, &organisation.LogoFileName,
			&organisation.CreatedAt, &organisation.UpdatedAt, &organisation.AccountHolderID,
			&organisationType.ID, &organisationType.Type, &organisationType.DisplayName)

	if err != nil {
		return models.Organisations{}, err
	}
	organisation.OrganisationType = &organisationType
	return organisation, nil
}

func (db *DB) GetSession(sessionToken string) (models.Session, error) {
	var session models.Session
	err := db.QueryRow(`
		SELECT id, user_id, session_token, ip_address, created_at, expires_at 
		FROM sessions 
		WHERE session_token = $1`, sessionToken).
		Scan(&session.ID, &session.UserID, &session.SessionToken, &session.IPAddress, &session.CreatedAt, &session.ExpiresAt)

	if err != nil {
		return models.Session{}, err
	}
	return session, nil
}

func (db *DB) CreateSession(sessionToken string, userID int, IpAddress string) error {
	_, err := db.Exec(`
        INSERT INTO sessions (user_id, session_token, ip_address, expires_at) 
        VALUES ($1, $2, $3, NOW() + INTERVAL '24 hours')`, userID, sessionToken, IpAddress)
	return err
}

func (db *DB) GetUserBySession(sessionToken string) (models.Users, error) {
	var user models.Users
	err := db.QueryRow(`
        SELECT u.id, u.email, u.password, u.organisation_id, u.name, u.created_at, u.updated_at 
        FROM users u
        INNER JOIN sessions s ON u.id = s.user_id 
        WHERE s.session_token = $1 AND s.expires_at > NOW()`, sessionToken).
		Scan(&user.ID, &user.Email, &user.Password, &user.OrganisationID, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Printf("Error getting user by session: %v", err)
		return models.Users{}, err
	}
	return user, nil
}

func (db *DB) GetBannedIp(ipAddress string) (models.BannedIPs, bool, error) {
	var bannedIp models.BannedIPs
	err := db.QueryRow(`
		SELECT id, ip_address, created_at, expires_at 
		FROM banned_ips 
		WHERE ip_address = $1 AND expires_at > NOW() AND expires_at < NOW() + INTERVAL '10 minutes'`, ipAddress).
		Scan(&bannedIp.ID, &bannedIp.IPAddress, &bannedIp.CreatedAt, &bannedIp.ExpiresAt)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return models.BannedIPs{}, false, nil
		}
		return models.BannedIPs{}, false, err
	}
	return bannedIp, true, nil
}

func (db *DB) CreateBannedIp(ipAddress string) error {

	_, err := db.Exec(`
		INSERT INTO banned_ips (ip_address, expires_at)
		VALUES ($1, NOW() + INTERVAL '10 minutes')`, ipAddress)

	if err != nil {
		return err
	}
	return nil
}

func (db *DB) CreateLoginAttempt(ipAddress string) (int, error) {
	_, err := db.Exec(`
		INSERT INTO login_attempts (ip_address) 
		VALUES ($1)
		`, ipAddress)

	if err != nil {
		return 0, err
	}

	var numberOfAttempts int

	err = db.QueryRow(`
		SELECT COUNT(*)
		FROM login_attempts
		WHERE ip_address = $1 AND created_at > NOW() - INTERVAL '1 hour'
		`, ipAddress).Scan(&numberOfAttempts)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return 0, nil
		}
		return 0, err
	}

	return numberOfAttempts, nil
}
