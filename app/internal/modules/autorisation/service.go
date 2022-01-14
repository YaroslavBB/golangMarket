package autorisation

import (
	"database/sql"
	"task/internal/entity/autorisatione"

	"github.com/jmoiron/sqlx"
)

type service struct {
	repo Repository
}

// NewService по autorisation
func NewService(r Repository) Service {
	return &service{
		r,
	}
}

func (s *service) SaveUser(tx *sqlx.Tx, user autorisatione.User) error {
	return s.repo.SaveUser(tx, user)
}

func (s *service) LoadUserByUsername(tx *sqlx.Tx, username string) (*autorisatione.User, error) {
	var userFromDB *autorisatione.User

	userFromDB, err := s.repo.LoadUserByUsername(tx, username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return userFromDB, nil
}
