package user

import (
	"context"
	"fmt"

	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateUser(ctx context.Context, req *desc.UserCreateRequest) (*desc.UserCreateResponse, error) {
	_, err := i.authorizer.AuthorizeAdmin(ctx)
	if err != nil {
		log.Error("unauthorized user",
			zap.Error(err),
		)
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized user")
	}

	err = i.validator.ValidateSecret(req.GetSecret())
	if err != nil {
		log.Error("validation error",
			zap.Error(err),
		)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("validation error: %v", err))
	}

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
