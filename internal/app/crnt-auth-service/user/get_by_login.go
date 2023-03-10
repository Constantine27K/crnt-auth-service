package user

import (
	"context"
	"database/sql"
	"errors"

	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetUserByLogin(ctx context.Context, req *desc.UserGetByLoginRequest) (*desc.UserGetByLoginResponse, error) {
	user, err := i.userStorage.GetByLogin(req.GetLogin())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Warn("no such user",
				zap.Any("request", req),
			)
			return nil, status.Error(codes.NotFound, "no such user")
		}
		log.Error("cannot get user by id",
			zap.String("login", req.GetLogin()),
			zap.Error(err))
		return nil, err
	}

	return &desc.UserGetByLoginResponse{User: user}, nil
}
