package autorization

import (
	"database/sql"
	"task/internal/entity/autorizatione"
	"task/internal/entity/global"

	"github.com/jmoiron/sqlx"
)

type repository struct {
}

// NewRepository репозиторий autorization
func NewRepository() Repository {
	return &repository{}
}

func (r *repository) SaveUser(tx *sqlx.Tx, user autorizatione.User) error {
	sqlQuery := `insert into users (username, password) values ($1, $2)`

	_, err := tx.Exec(sqlQuery, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) LoadUserByUsername(tx *sqlx.Tx, username string) (*autorizatione.User, error) {
	sqlQuery := `select u.username, u.password from users as u where username = $1`

	var user autorizatione.User

	err := tx.Get(&user, sqlQuery, username)
	switch err {
	case nil:
		return &user, nil
	case sql.ErrNoRows:
		return nil, global.ErrNoDataFound
	default:
		return nil, err
	}
}
