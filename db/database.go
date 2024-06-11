// Package db provides database access.
package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

var pool *pgxpool.Pool

// Migrate runs all outstanding database migrations.
func Migrate() error {
	var err error
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}

	err = validateSchema(db)
	if err != nil {
		return err
	}

	return nil
}

type queryTracer struct{}

func (qt *queryTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	slog.Info("Query", "sql", data.SQL, "args", data.Args)
	return ctx
}

func (qt *queryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
}

// GetPool returns a database connection pool.
func GetPool() (*pgxpool.Pool, error) {
	if pool == nil {
		pgConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse DATABASE_URL: %v\n", err)
			return nil, err
		}

		// pgConfig.ConnConfig.Tracer = &queryTracer{}

		pgConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
			pgxUUID.Register(conn.TypeMap())
			return nil
		}

		pool, err = pgxpool.NewWithConfig(context.Background(), pgConfig)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			return nil, err
		}
	}

	return pool, nil
}
