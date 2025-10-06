package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", mainRoute)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}

func mainRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", rand.IntN(6)+1)
}
