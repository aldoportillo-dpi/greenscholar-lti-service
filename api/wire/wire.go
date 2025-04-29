//go:build wireinject
// +build wireinject

package wire

import (
	"log/slog"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/illinoisdpi/go-service-template/api/users"
	"github.com/illinoisdpi/go-service-template/configs"
	"github.com/illinoisdpi/go-service-template/db"
	"github.com/illinoisdpi/go-service-template/db/sqlc"
)

func InitializeLogger() *slog.Logger {
	wire.Build(configs.NewLogger)
	return &slog.Logger{}
}

func InitializeDBPool() *pgxpool.Pool {
	wire.Build(db.NewDBPool)
	return nil
}

func InitializeDBQueries() *sqlc.Queries {
	wire.Build(db.NewQueries, InitializeDBPool)
	return nil
}

func InitializeUserService() *users.UserService {
	wire.Build(users.NewUserService, InitializeLogger, InitializeDBQueries)
	return &users.UserService{}
}
