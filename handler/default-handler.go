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

func CreateItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("request for create item: " + id))
}

func UpdateItemById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("request for update item: " + id))
}

func DeleteItemById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("request for delete item: " + id))
}
