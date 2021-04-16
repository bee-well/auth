package domain

import (
	"database/sql"

	"github.com/bee-well/auth/config"
)

type SqlInterface interface {
	Connect() (*sql.DB, error)
}

type postgreSql struct{}

func newSqlConnector() SqlInterface {
	return &postgreSql{}
}

func (postgreSql) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.SqlConnectionUrl())
	if err != nil {
		return nil, err
	}
	return db, nil
}
