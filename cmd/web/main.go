package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"snippetbox.anandgautam.io/internal/models"
)

func main() {

	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.dsn, "dsn", "web:123456@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// To keep the main() function tiday, put the code for creating a new
	//connection pool into a separate function named openDB(). We pass openDB()
	//the DSN from the command-line flag, and it returns a sql.DB connection pool.

	db, err := openDB(cfg.dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Defer a call to db.Close() so that the connection pool is closed before
	//the main() function exits.
	defer db.Close()

	// Intialize a new instance of our application struct, containing the
	//dependencies
	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("Starting server", slog.String("addr", cfg.addr))

	err = http.ListenAndServe(cfg.addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}

// openDB() function
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Use the Ping() method to check that the connection is working as expected.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
