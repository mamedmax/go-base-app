package handler

import "net/http"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello 👋🏻"))
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is default handler"))
}
