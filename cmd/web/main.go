// Package main is the entry point for the web server.
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"example.com/webserver/internal/app/boot"
	"example.com/webserver/internal/app/config"
)

func main() {
	cfg, notices, configErr := config.Load()
	if configErr != nil {
		panic(configErr)
	}

	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	if cfg.Env == "dev" {
		opts.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	for _, msg := range notices {
		logger.Warn(msg)
	}

	app, appErr := boot.New(logger, cfg)
	if appErr != nil {
		panic(appErr)
	}

	src := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.PORT),
		Handler:           app.Router,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		logger.Info("server started",
			slog.Int("port", cfg.PORT),
			slog.String("env", cfg.Env),
		)
		errCh <- src.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("something went wrong", slog.Any("error", err))
			return
		}
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		if err := src.Shutdown(shutdownCtx); err != nil {
			logger.Error("server forced to shutdown", slog.Any("error", err))
			return
		}
	}

	logger.Info("server shutdown")
}
