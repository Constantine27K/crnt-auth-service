package user

import (
	"context"

	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateUser(ctx context.Context, req *desc.UserUpdateRequest) (*desc.UserUpdateResponse, error) {
	payload, err := i.authorizer.AuthorizeUser(ctx)
	if err != nil {
		log.Error("unauthorized user",
			zap.Error(err),
		)
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized user")
	}

	if payload.Username != req.GetUser().GetDisplayName() {
		return nil, status.Errorf(codes.PermissionDenied, "you can edit only your user")
	}

	id, err := i.userStorage.Update(req.GetId(), req.GetUser())
	if err != nil {
		log.Error("cannot update user",
			zap.Int64("id", req.GetId()),
			zap.Error(err),
		)
		return nil, err
	}

	return &desc.UserUpdateResponse{Id: id}, nil
}
