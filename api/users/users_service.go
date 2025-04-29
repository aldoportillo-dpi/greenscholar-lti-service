package users

import (
	"log/slog"

	"github.com/illinoisdpi/go-service-template/db/sqlc"
)

type UserService struct {
	logger  *slog.Logger
	queries *sqlc.Queries
}

func NewUserService(logger *slog.Logger, queries *sqlc.Queries) *UserService {
	return &UserService{
		logger,
		queries,
	}
}
