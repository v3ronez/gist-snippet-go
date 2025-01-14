package main

import (
	"log"
	"net/http"
)

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed XD"))
		return
	}
	w.Write([]byte("create snippet"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("view"))
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return

	}
	w.Write([]byte("lets go"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snipppet", snippetView)
	mux.HandleFunc("POST /snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Server is running on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatalln(err)
}
