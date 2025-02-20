package api

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gerry-sheva/bts-todo-list/pkg/auth"
	"github.com/gerry-sheva/bts-todo-list/pkg/checklist"
	"github.com/jackc/pgx/v5/pgxpool"
)

type app struct {
	port   int
	logger *slog.Logger
	dbpool *pgxpool.Pool
}

// Create and start server api
func StartServer(logger *slog.Logger, dbpool *pgxpool.Pool) {
	// Config for port
	var port int
	flag.IntVar(&port, "port", 8080, "API server port")
	flag.Parse()

	app := &app{
		port,
		logger,
		dbpool,
	}

	AuthHandler := auth.New(app.logger, app.dbpool)
	ChecklistHandler := checklist.New(app.logger, app.dbpool)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", AuthHandler.RegisterUser)
	mux.HandleFunc("POST /login", AuthHandler.LoginUser)
	mux.Handle("POST /checklist", Auth(http.HandlerFunc(ChecklistHandler.CreateChecklist)))
	mux.Handle("DELETE /checklist/{checklist_id}", Auth(http.HandlerFunc(ChecklistHandler.DeleteChecklist)))
	mux.Handle("GET /checklist", Auth(http.HandlerFunc(ChecklistHandler.GetAllChecklist)))
	mux.Handle("GET /checklist/{checklist_id}", Auth(http.HandlerFunc(ChecklistHandler.GetChecklistDetails)))

	muxWithMiddleware := LogRequests(logger)(mux)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.port),
		Handler:      muxWithMiddleware,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		logger.Info("shutting down server", slog.String("signal", s.String()))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx)
	}()

	logger.Info("starting server", "port", app.port)

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("server error", "error", err)
		os.Exit(1)
	}

	shutdownErr := <-shutdownError
	if shutdownErr != nil {
		logger.Error("graceful shutdown failed", "error", shutdownErr)
	} else {
		logger.Info("stopped server", slog.String("addr", srv.Addr))
	}
}
