package repository

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type UsersRepository interface {
	CountUserByUsername(username string) (int, error)
	CreateUser(u User) error
	GetUser(username string) (User, error)
}

type usersRepository struct{}

func NewUsersRepository() UsersRepository {
	return usersRepository{}
}
