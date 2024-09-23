package handler

import "net/http"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello 👋🏻"))
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is default handler"))
}

func FindItemById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("request for item: " + id))
}
