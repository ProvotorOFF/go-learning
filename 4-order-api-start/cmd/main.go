package main

import (
	"net/http"
	"order-api-start/configs"
	"order-api-start/internal/auth"
	"order-api-start/internal/product"
	"order-api-start/internal/session"
	"order-api-start/internal/user"
	"order-api-start/pkg/db"
	"order-api-start/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	db, _ := db.NewDb(conf)

	//Repos
	productRepository := product.NewProductRepository(db)
	userRepository := user.NewUserRepository(db)
	sessionRepo := session.NewSessionRepository(db)

	//Services
	authService := auth.NewAuthService(auth.ServiceDeps{
		UserRepo:    userRepository,
		SessionRepo: sessionRepo,
		Config:      conf,
	})

	//Handlers
	product.NewProductHandler(router, product.Deps{
		Repo: productRepository,
	})
	auth.NewAuthHandler(router, auth.Deps{
		Repo:    userRepository,
		Service: authService,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logger(router),
	}

	server.ListenAndServe()
}
