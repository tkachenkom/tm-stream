package cfg

import (
	"database/sql"

	// blank import for init pq driver
	_ "github.com/lib/pq"
)

type Config struct {
	DB *sql.DB
}

func New() (*Config, error) {
	dbConf := getDB()

	err := dbConf.Validate()
	if err != nil {
		return nil, err
	}

	uriDB := dbConf.linkWrap()
	conn, err := sql.Open("postgres", uriDB)
	if err != nil {
		return nil, err
	}

	return &Config{
		DB: conn,
	}, nil
}
