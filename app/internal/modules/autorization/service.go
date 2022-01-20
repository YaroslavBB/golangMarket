package autorization

import (
	"task/internal/entity/autorizatione"

	"github.com/jmoiron/sqlx"
)

type service struct {
	repo Repository
}

// NewService по autorization
func NewService(r Repository) Service {
	return &service{
		r,
	}
}

func (s *service) SaveUser(tx *sqlx.Tx, user autorizatione.User) error {
	return s.repo.SaveUser(tx, user)
}

func (s *service) LoadUserByUsername(tx *sqlx.Tx, username string) (*autorizatione.User, error) {
	return s.repo.LoadUserByUsername(tx, username)
}
