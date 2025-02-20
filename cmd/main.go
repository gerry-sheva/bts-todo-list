package main

import (
	"log/slog"
	"os"

	"github.com/gerry-sheva/bts-todo-list/pkg/api"
	"github.com/gerry-sheva/bts-todo-list/pkg/database"
	"github.com/joho/godotenv"
)

// Entry point of the application
// Initializes dependencies and start the server
func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	err := godotenv.Load()
	if err != nil {
		logger.Error("Failed to load env file")
		return
	}

	dbpool := database.ConnectDB()
	api.StartServer(logger, dbpool)
}
