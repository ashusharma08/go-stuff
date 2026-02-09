package main

import (
	"log/slog"
	"os"
)

func main() {
	InitLogger("dev")
	slog.Info("App starting")
}

func InitLogger(env string) {
	var handler slog.Handler

	if env == "prod" {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = slog.NewTextHandler(os.Stdout, nil)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
