package repository

import "log"

func (User) CountUserByUsername(username string) (int, error) {
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM users WHERE username = ?", username)
	if err != nil {
		log.Printf("Error counting user by username: %v", err)
		return 0, err
	}

	return count, nil
}

func (User) CreateUser(u User) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", u.Username, u.Password)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}
