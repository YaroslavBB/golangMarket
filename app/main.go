package main

import (
	"log"
	"task/config"
	"task/internal/app/rest"
	"task/internal/modules/product"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	confPath := "config/config.yaml"
	psqlInfo := config.GetConfiguration(confPath)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	r.Use(cors.Default())

	productRepo := product.NewRepository()
	productService := product.NewService(productRepo)

	server := rest.NewRest(r, db, productService)

	server.Run()
}