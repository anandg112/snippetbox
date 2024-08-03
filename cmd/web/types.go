package main

import "log/slog"

type config struct {
	addr      string
	staticDir string
}

type application struct {
	logger *slog.Logger
}
