package main

import (
	"fmt"
	"go-base-app/handler"
	"go-base-app/middleware"
	"net/http"
)

func main() {

	reqHandler := handler.NewHandler()

	router := http.NewServeMux()
	router.HandleFunc("/", reqHandler.HelloHandler)
	router.HandleFunc("GET /item/{id}", reqHandler.FindItemById)
	router.HandleFunc("POST /item/{id}", reqHandler.CreateItem)
	router.HandleFunc("PUT /item/{id}", reqHandler.UpdateItemById)
	router.HandleFunc("DELETE /item/{id}", reqHandler.DeleteItemById)

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1/", router))

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("Server listening on port :8080")
	server.ListenAndServe()
}
