package repository

import "log"

func (ur usersRepository) CountUserByUsername(username string) (int, error) {
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM users WHERE username = ?", username)
	if err != nil {
		log.Printf("Error counting user by username: %v", err)
		return 0, err
	}

	return count, nil
}

func (ur usersRepository) CreateUser(u User) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", u.Username, u.Password)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

func (ur usersRepository) GetUser(username string) (User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return User{}, err
	}
	return user, nil
}
