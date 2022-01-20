package autorization

import (
	"task/internal/entity/autorizatione"

	"github.com/jmoiron/sqlx"
)

// Repository по autorization
type Repository interface {
	SaveUser(tx *sqlx.Tx, user autorizatione.User) error
	LoadUserByUsername(tx *sqlx.Tx, username string) (*autorizatione.User, error)
}

// Service по autorization
type Service interface {
	SaveUser(tx *sqlx.Tx, user autorizatione.User) error
	LoadUserByUsername(tx *sqlx.Tx, username string) (*autorizatione.User, error)
}
