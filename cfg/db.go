package cfg

import (
	"fmt"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	dbHost     = "HOST_DB"
	dbPort     = "PORT_DB"
	dbUser     = "USER_DB"
	dbPassowrd = "PASSWORD_DB"
	dbName     = "NAME_DB"
	dbSSL      = "SSL_DB"
)

type DB struct {
	Host     string `env:"HOST_DB"`
	Port     string `env:"PORT_DB"`
	User     string `env:"USER_DB"`
	Password string `env:"PASSWORD_DB"`
	Name     string `env:"NAME_DB"`
	SSL      string `env:"SSL_DB"`
}

func (d *DB) Validate() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.Host, validation.Required),
		validation.Field(&d.Port, validation.Required),
		validation.Field(&d.User, validation.Required),
		validation.Field(&d.Password, validation.Required),
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.SSL, validation.Required),
	)
}

func getDB() *DB {
	return &DB{
		Host:     os.Getenv(dbHost),
		Port:     os.Getenv(dbPort),
		User:     os.Getenv(dbUser),
		Password: os.Getenv(dbPassowrd),
		Name:     os.Getenv(dbName),
		SSL:      os.Getenv(dbSSL),
	}
}

func (d *DB) linkWrap() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s", d.Host, d.Port, d.User, d.Password, d.Name, d.SSL)
}
