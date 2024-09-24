package repository

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/srinathgs/mysqlstore"
)

var db *sqlx.DB

func NewConnection(conf mysql.Config) error {
	_db, err := sqlx.Open("mysql", conf.FormatDSN())
	if err != nil {
		return err
	}

	db = _db
	return nil
}

func NewStore() (*mysqlstore.MySQLStore, error) {
	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "sessions", "/", 60*60*24*14, []byte("secret"))
	if err != nil {
		return nil, err
	}

	return store, nil
}
