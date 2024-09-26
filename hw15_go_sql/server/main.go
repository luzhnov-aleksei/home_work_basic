package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luzhnov-aleksei/home_work_basic/hw15_go_sql/internal/config"
	"github.com/luzhnov-aleksei/home_work_basic/hw15_go_sql/internal/http-server/handlers"
	mwLogger "github.com/luzhnov-aleksei/home_work_basic/hw15_go_sql/internal/http-server/middleware/logger"
	"github.com/luzhnov-aleksei/home_work_basic/hw15_go_sql/internal/lib/logger/sl"
	"github.com/luzhnov-aleksei/home_work_basic/hw15_go_sql/internal/storage"
	"golang.org/x/exp/slog"
)

var Storage storage.Storage

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("Starting online-shop", slog.String("env", cfg.Env))
	log.Debug("DEBUG messages are enabled", slog.String("env", cfg.Env))

	err := Storage.NewDB(cfg.Driver, cfg.Dsn)
	if err != nil {
		log.Error("Failed to init database", sl.Err(err))
		os.Exit(1)
	}

	err = Storage.Ping()
	if err != nil {
		log.Error("Failed to ping database", sl.Err(err))
		os.Exit(1)
	}

	defer Storage.CloseDB()

	log.Info("Successfully connected!", slog.String("driver", cfg.Driver), slog.String("dsn", cfg.Dsn))

	router := chi.NewRouter()

	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.HandleFunc("/users", handlers.UsersHandler(log, &Storage))
	router.HandleFunc("/users/stat", handlers.UserStatHandler(log, &Storage))
	router.HandleFunc("/products", handlers.ProductsHandler(log, &Storage))
	router.HandleFunc("/orders", handlers.OrdersHandler(log, &Storage))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	log.Info("Starting server", slog.String("address", srv.Addr))

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Error("Failed to serve server", sl.Err(err))
		}
	}()

	log.Info("Server started")

	<-done
	log.Info("Stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.HTTPServer.ShutdownTimeout)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Error("Failed to stop server", sl.Err(err))
		return
	}

	log.Info("Server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
