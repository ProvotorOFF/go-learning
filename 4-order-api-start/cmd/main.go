package main

import (
	"net/http"
	"order-api-start/configs"
	"order-api-start/internal/auth"
	"order-api-start/internal/order"
	"order-api-start/internal/product"
	"order-api-start/internal/session"
	"order-api-start/internal/user"
	"order-api-start/pkg/db"
	"order-api-start/pkg/middleware"
)

func main() {

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logger(App()),
	}

	server.ListenAndServe()
}

func App() http.Handler {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	db, _ := db.NewDb(conf)

	//Repos
	productRepository := product.NewProductRepository(db)
	userRepository := user.NewUserRepository(db)
	sessionRepo := session.NewSessionRepository(db)
	orderRepo := order.NewOrderRepository(db)

	//Services
	authService := auth.NewAuthService(auth.ServiceDeps{
		UserRepo:    userRepository,
		SessionRepo: sessionRepo,
		Config:      conf,
	})

	//Handlers
	product.NewProductHandler(router, product.Deps{
		Repo: productRepository,
		Conf: conf,
	})
	auth.NewAuthHandler(router, auth.Deps{
		Repo:    userRepository,
		Service: authService,
	})
	order.NewOrderHandler(router, order.Deps{
		Repo:     orderRepo,
		Conf:     conf,
		UserRepo: userRepository,
	})

	return router
}
