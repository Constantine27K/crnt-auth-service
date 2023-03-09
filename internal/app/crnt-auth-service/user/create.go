package user

import (
	"context"

	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) CreateUser(ctx context.Context, req *desc.UserCreateRequest) (*desc.UserCreateResponse, error) {
	secretID, err := i.secretStorage.Add(req.GetSecret())
	if err != nil {
		log.Error("cannot store secret",
			zap.Error(err),
		)
		return nil, err
	}

	userID, err := i.userStorage.Add(req.GetUser(), secretID)
	if err != nil {
		log.Error("cannot store user",
			zap.Any("user", req.GetUser()),
			zap.Error(err),
		)
		return nil, err
	}

	return &desc.UserCreateResponse{Id: userID}, nil
}
