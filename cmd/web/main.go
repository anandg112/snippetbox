package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Intialize a new instance of our application struct, containing the
	//dependencies
	app := &application{
		logger: logger,
	}

	logger.Info("Starting server", slog.String("addr", cfg.addr))

	err := http.ListenAndServe(cfg.addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}
