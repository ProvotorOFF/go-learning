package main

import (
	"net/http"
	"order-api-start/configs"
	"order-api-start/internal/product"
	"order-api-start/pkg/db"
	"order-api-start/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	db, _ := db.NewDb(conf)

	//Repos
	productRepository := product.NewProductRepository(db)

	//Handlers
	product.NewProductHandler(router, product.Deps{
		Repo: productRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logger(router),
	}

	server.ListenAndServe()
}
