package main

import (
	"net/http"
	"validation-api/configs"
	"validation-api/internal/verify"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()

	verify.NewVerifyHandler(router, conf)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
