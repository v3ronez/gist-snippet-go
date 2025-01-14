package main

import (
	"log"
	"net/http"
)

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create snippet"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("view"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("lets go"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snipppet", snippetView)

	err := http.ListenAndServe(":4000", mux)
	log.Fatalln(err)
}
