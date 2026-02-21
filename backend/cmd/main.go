package main

import (
	"log"
	"log/slog"
	"os"
	"simple-todo/helpers"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := helpers.InitStorage("data"); err != nil {
		log.Fatalf("Failed to create base directory: %v", err)
	}
}
