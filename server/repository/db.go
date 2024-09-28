package repository

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/srinathgs/mysqlstore"
)

var (
	db       *sqlx.DB
	userName map[int]string
	userID   map[string]int
)

func NewConnection(conf mysql.Config) error {
	_db, err := sqlx.Open("mysql", conf.FormatDSN())
	if err != nil {
		return err
	}

	db = _db

	userName = make(map[int]string)
	userID = make(map[string]int)
	return nil
}

func NewStore() (*mysqlstore.MySQLStore, error) {
	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "session", "/", 60*60*24*14, []byte("secret"))
	if err != nil {
		return nil, err
	}

	return store, nil
}
