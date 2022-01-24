package main

import (
	"log"
	"os"
	"task/config"
	"task/internal/app/rest"
	"task/internal/modules/autorization"
	"task/internal/modules/product"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	config := config.NewConfig(os.Getenv("CONF_PATH"))

	db, err := sqlx.Open("postgres", config.GetConfiguration())
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	productRepo := product.NewRepository()
	productService := product.NewService(productRepo)

	autorizationRepo := autorization.NewRepository()
	autorizationService := autorization.NewService(autorizationRepo)

	server := rest.NewRest(r, db, productService, autorizationService)

	server.Run()
}
