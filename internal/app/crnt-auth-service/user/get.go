package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/users/models"
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetUsers(ctx context.Context, req *desc.UserGetRequest) (*desc.UserGetResponse, error) {
	filter := models.NewUsersFilter(req)

	users, err := i.userStorage.Get(filter)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Warn("no such user",
				zap.Any("request", req),
			)
			return nil, status.Error(codes.NotFound, "no such user")
		}
		log.Error("cannot get users",
			zap.Any("request", req),
			zap.Error(err),
		)
		return nil, err
	}

	return &desc.UserGetResponse{Users: users}, nil
}
