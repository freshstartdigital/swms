package repository

import "example.com/internal/models"

func (db *DB) LoginHandler(email string, password string) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT * FROM users WHERE email = $1 AND password = $2", email, password).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *DB) StoreSessionTokenHandler(sessionToken string, userID int) error {
	_, err := db.Exec("INSERT INTO sessions (user_id, session_token, expires_at) VALUES ($1, $2, NOW() + INTERVAL '24 hours')", userID, sessionToken)
	if err != nil {
		return err
	}
	return nil

}

func (db *DB) AuthenticateHandler(sessionToken string) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT users.id, users.email, users.password, users.created_at, users.updated_at FROM users INNER JOIN sessions ON users.id = sessions.user_id WHERE sessions.session_token = $1 AND sessions.expires_at > NOW()", sessionToken).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
