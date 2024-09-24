package repository

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

type UsersRepository interface {
	CountUserByUsername(username string) (int, error)
	CreateUser(u User) error
	GetUser(username string) (User, error)
}
