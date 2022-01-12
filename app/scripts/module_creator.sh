#!/bin/bash

##!/bin/bash

if [ -z "$1" ]
  then
    echo "Укажите название модуля"
    exit 1
fi


MODULE=$1
MODULE_DIR=../internal/modules/$1

echo "creating the module ${MODULE}"
mkdir -p $MODULE_DIR

echo "package ${MODULE} 

// Repository по ${MODULE} 
type Repository interface {
}

// Service по ${MODULE}
type Service interface {
}" > $MODULE_DIR/interface.go

echo "package ${MODULE}

type repository struct {
}
// NewRepository репозиторий ${MODULE}
func NewRepository() Repository {
	return &repository{}
}
" > $MODULE_DIR/repository_oracle.go

echo "package ${MODULE}_test

const (
	confPath = \"../../../config/conf.yaml\"
)

// func TestIsNewClient(t *testing.T) {
// 	r := require.New(t)

// 	config, err := config.GetProjectConf(confPath)
// 	r.NoError(err)

// 	db := db.GetSqlxDB(config.OracleConnectString())
// 	defer db.Close()

// 	tx, err := db.Beginx()
// 	r.NoError(err)
// 	defer tx.Rollback()

// 	repo := tariffdata.NewRepository(db)

// 	const (
// 		newClient = 85467
// 		oldClient = 3378
// 	)

// 	t.Run(\"новый клиент\", func(t *testing.T) {
// 		new, err := repo.IsNewClient(tx, newClient)
// 		r.NoError(err)
// 		r.True(new)
// 	})

// 	t.Run(\"старый клиент\", func(t *testing.T) {
// 		new, err := repo.IsNewClient(tx, newClient)
// 		r.NoError(err)
// 		r.False(new)
// 	})

// }


" > $MODULE_DIR/repository_test.go
echo "package ${MODULE}

type service struct {
	repo Repository
}

// NewService по ${MODULE}
func NewService(r Repository) Service {
	return &service{
		r,
	}
}
" > $MODULE_DIR/service.go
echo "package ${MODULE}_test" > $MODULE_DIR/service_test.go
mkdir -p ../internal/entity/${MODULE}e
echo "package ${MODULE}e" > ../internal/entity/${MODULE}e/${MODULE}e.go
echo "package ${MODULE}e" > ../internal/entity/${MODULE}e/errors.go
echo "package ${MODULE}e" > ../internal/entity/${MODULE}e/constants.go