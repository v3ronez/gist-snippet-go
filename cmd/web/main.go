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
	config   *config
}

func initApp() *application {
	config := &config{
		addr:      "",
		staticDir: "",
	}

	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return &application{
		infoLog:  logInfo,
		errorLog: errorLog,
		config:   config,
	}
}

func main() {
	app := initApp()

	flag.StringVar(&app.config.addr, "addr", ":4000", "Http Network address")
	flag.StringVar(&app.config.staticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	// f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer f.Close()
	//
	// infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)

	app.infoLog.Printf("Server running on %s\n", app.config.addr)
	slog.Info("Server configs:", "addr", app.config.addr, "static-dir", app.config.staticDir)

	srv := &http.Server{
		Addr:     app.config.addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		app.errorLog.Fatal(err)
	}
}
