package main

import (
	"log/slog"

	"snippetbox.anandgautam.io/internal/models"
)

type config struct {
	addr string
	dsn  string
}

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}
