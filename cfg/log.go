package cfg

import "github.com/go-ozzo/ozzo-validation"

type Log struct {
	Level string `env:"LOG_LEVEL"`
}

func (l *Log) Validate() error {
	return validation.ValidateStruct(l,
		validation.Field(&l.Level, validation.Required),
	)
}
