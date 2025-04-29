package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/illinoisdpi/go-service-template/db/sqlc"
)

var poolConn *pgxpool.Pool

func NewDBPool() *pgxpool.Pool {
	if poolConn != nil {
		return poolConn
	}

	var err error
	poolConn, err = pgxpool.New(context.Background(), os.Getenv("CUSTOMCONNSTR_DB_CONN_STR"))
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	return poolConn
}

func RunMigrations() error {
	m, err := migrate.New(
		"file://db/migrations",
		os.Getenv("CUSTOMCONNSTR_DB_CONN_STR"))
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}

		return err
	}

	return nil
}

func NewQueries(db *pgxpool.Pool) *sqlc.Queries {
	return sqlc.New(db)
}
