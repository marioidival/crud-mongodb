package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, path %s!", r.URL.Path[1:])
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
