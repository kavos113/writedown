package repository

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(conf mysql.Config) (Repository, error) {
	db, err := sqlx.Open("mysql", conf.FormatDSN())
	if err != nil {
		return Repository{}, err
	}

	return Repository{db: db}, nil
}
