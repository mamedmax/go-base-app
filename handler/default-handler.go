package handler

import (
	"fmt"
	"go-base-app/middleware"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello 👋🏻"))
}

func (h *Handler) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is default handler"))
}

func (h *Handler) FindItemById(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.AuthUserID).(string)
	if !ok {
		fmt.Println("Invalid user ID")
		w.WriteHeader(http.StatusBadRequest)
	}

	id := r.PathValue("id")
	fmt.Println("user ID: " + userID)
	w.Write([]byte("request for item: " + id + " from userID: " + userID))
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("request for create item: " + id))
}

func (h *Handler) UpdateItemById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("request for update item: " + id))
}

func (h *Handler) DeleteItemById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("request for delete item: " + id))
}
