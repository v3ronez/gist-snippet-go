package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
)

type config struct {
	addr      string
	staticDir string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "Http Network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	fileStaticServer := http.FileServer(http.Dir(cfg.staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileStaticServer))

	log.Printf("Server running on %s\n", cfg.addr)
	slog.Info("Server configs:", "addr", cfg.addr, "static-dir", cfg.staticDir)

	err := http.ListenAndServe(cfg.addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
