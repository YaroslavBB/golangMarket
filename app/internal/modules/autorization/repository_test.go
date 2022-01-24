package autorization_test

import (
	"task/config"
	"task/internal/entity/autorizatione"
	"task/internal/modules/autorization"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

const (
	confPath = "../../../config/config.yaml"
)

var testUser = autorizatione.User{Username: "userTestRepo", Password: "user1"}

func TestSaveUser(t *testing.T) {
	config := config.NewConfig(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config.GetConfiguration())
	require.NoError(t, err)
	defer db.Close()

	repo := autorization.NewRepository()

	t.Run("добавление юзера", func(t *testing.T) {
		tx, err := db.Beginx()
		require.NoError(t, err)
		defer tx.Rollback()

		err = repo.SaveUser(tx, testUser)
		require.NoError(t, err)

		t.Run("Проверка", func(t *testing.T) {
			var users []autorizatione.User

			err := tx.Select(&users, `select u.username, u.password from users as u`)
			require.NoError(t, err)

			require.Contains(t, users, testUser)
		})
	})
}

func TestLoadUserByUsername(t *testing.T) {

	config := config.NewConfig(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config.GetConfiguration())
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := autorization.NewRepository()

	t.Run("добавление юзера", func(t *testing.T) {
		err := repo.SaveUser(tx, testUser)
		require.NoError(t, err)

		t.Run("проверка поиска по имени", func(t *testing.T) {
			var userFromDB *autorizatione.User

			userFromDB, err = repo.LoadUserByUsername(tx, testUser.Username)
			require.NoError(t, err)

			require.Equal(t, testUser.Username, userFromDB.Username)
		})
	})
}
