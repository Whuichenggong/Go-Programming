package main

import (
	"fmt"
	"net/http"
)

//1.22版本最新特
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "renturna single comment for comment id: %s", id)
	})

	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "POST a new commment")
	})

	http.ListenAndServe(":8080", mux)
}
