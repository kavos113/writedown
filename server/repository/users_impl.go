package repository

import "log"

func (r Repository) CountUserByUsername(username string) (int, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM users WHERE username = ?", username)
	if err != nil {
		log.Printf("Error counting user by username: %v", err)
		return 0, err
	}

	return count, nil
}

func (r Repository) CreateUser(u User) error {
	_, err := r.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", u.Username, u.Password)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}
