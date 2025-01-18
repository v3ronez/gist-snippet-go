package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func initApp() *application {

	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return &application{
		infoLog:  logInfo,
		errorLog: errorLog,
	}
}

func main() {
	var cfg config
	app := initApp()

	flag.StringVar(&cfg.addr, "addr", ":4000", "Http Network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	fileStaticServer := http.FileServer(http.Dir(cfg.staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileStaticServer))

	// f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer f.Close()
	//
	// infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)

	app.infoLog.Printf("Server running on %s\n", cfg.addr)
	slog.Info("Server configs:", "addr", cfg.addr, "static-dir", cfg.staticDir)

	srv := &http.Server{
		Addr:     *&cfg.addr,
		ErrorLog: app.errorLog,
		Handler:  mux,
	}

	err := srv.ListenAndServe()

	if err != nil {
		app.errorLog.Fatal(err)
	}
}
