package main

import (
	"fmt"
	"go-base-app/handler"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", handler.HelloHandler)
	router.HandleFunc("GET /item/{id}", handler.FindItemById)
	router.HandleFunc("POST /item/{id}", handler.CreateItem)
	router.HandleFunc("PUT /item/{id}", handler.UpdateItemById)
	router.HandleFunc("DELETE /item/{id}", handler.DeleteItemById)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Starting server on port :8080")
	server.ListenAndServe()
}
