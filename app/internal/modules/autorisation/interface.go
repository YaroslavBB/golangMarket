package autorisation

import (
	"task/internal/entity/autorisatione"

	"github.com/jmoiron/sqlx"
)

// Repository по autorisation
type Repository interface {
	SaveUser(tx *sqlx.Tx, user autorisatione.User) error
	LoadUserByUsername(tx *sqlx.Tx, username string) (*autorisatione.User, error)
}

// Service по autorisation
type Service interface {
	SaveUser(tx *sqlx.Tx, user autorisatione.User) error
	LoadUserByUsername(tx *sqlx.Tx, username string) (*autorisatione.User, error)
}
