package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"simple-todo/helpers"
	"simple-todo/internal/route"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := helpers.InitStorage("data"); err != nil {
		log.Fatalf("Failed to create base directory: %v", err)
	}

	router := route.CreateRoute()
	slog.Info("Starting server port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
