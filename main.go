package main

import (
	"fmt"
	"go-base-app/handler"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", handler.HelloHandler)
	router.HandleFunc("/item/{id}", handler.FindItemById)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Starting server on port :8080")
	server.ListenAndServe()
}
