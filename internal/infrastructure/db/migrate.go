package db

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log/slog"

	"example.com/webserver/internal/app/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/pressly/goose/v3/lock"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

//go:embed seeds/*.sql
var seedFiles embed.FS

// RunMigrations executes the goose migrations.
func RunMigrations(ctx context.Context, dbPool *pgxpool.Pool, logger *slog.Logger) error {
	provider, provErr := setupProvider(dbPool, logger, migrationFiles, "migrations")
	if provErr != nil {
		return fmt.Errorf("could not create goose provider: %w", provErr)
	}
	defer func() {
		if err := provider.Close(); err != nil {
			logger.Error("could not close provider", slog.Any("error", err))
		}
	}()

	if err := provider.Ping(ctx); err != nil {
		return fmt.Errorf("could not ping database with goose provider: %w", err)
	}

	_, migErr := provider.Up(ctx)
	if migErr != nil {
		return fmt.Errorf("could not run migrations: %w", migErr)
	}

	return nil
}

// RunSeeds executes the goose seed migrations.
func RunSeeds(ctx context.Context, dbPool *pgxpool.Pool, logger *slog.Logger, cfg *config.Config) error {
	if cfg.IsProduction() {
		return nil
	}

	provider, provErr := setupProvider(dbPool, logger, seedFiles, "seeds", goose.WithTableName("goose_db_seeds"))
	if provErr != nil {
		return fmt.Errorf("could not create goose provider: %w", provErr)
	}
	defer func() {
		if err := provider.Close(); err != nil {
			logger.Error("could not close provider", slog.Any("error", err))
		}
	}()

	if err := provider.Ping(ctx); err != nil {
		return fmt.Errorf("could not ping database with goose provider: %w", err)
	}

	_, migErr := provider.Up(ctx)
	if migErr != nil {
		return fmt.Errorf("could not run seeds: %w", migErr)
	}

	return nil
}

func setupProvider(
	dbPool *pgxpool.Pool,
	logger *slog.Logger,
	files embed.FS,
	filesSubDirName string,
	opts ...goose.ProviderOption,
) (*goose.Provider, error) {
	db := stdlib.OpenDBFromPool(dbPool)

	locker, lockErr := lock.NewPostgresSessionLocker()
	if lockErr != nil {
		return nil, fmt.Errorf("could not create goose session locker: %w", lockErr)
	}

	filesSub, fsErr := fs.Sub(files, filesSubDirName)
	if fsErr != nil {
		return nil, fmt.Errorf("could not create sub filesystem for migrations: %w", fsErr)
	}

	opts = append(opts,
		goose.WithVerbose(true),
		goose.WithSessionLocker(locker),
		goose.WithSlog(logger),
	)

	return goose.NewProvider(
		goose.DialectPostgres,
		db,
		filesSub,
		opts...,
	)
}
