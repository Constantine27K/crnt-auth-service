package user

import (
	"context"

	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateUserContacts(ctx context.Context, req *desc.UserContactsUpdateRequest) (*desc.UserContactsUpdateResponse, error) {
	payload, err := i.authorizer.AuthorizeUser(ctx)
	if err != nil {
		log.Error("unauthorized user",
			zap.Error(err),
		)
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized user")
	}

	user, err := i.userStorage.GetByID(req.GetId())
	if err != nil {
		log.Error("cannot get user by id",
			zap.Int64("id", req.GetId()),
			zap.Error(err),
		)
		return nil, err
	}

	if payload.Username != user.GetDisplayName() {
		return nil, status.Errorf(codes.PermissionDenied, "you can edit only your user")
	}

	id, err := i.userStorage.UpdateContacts(req.GetId(), req.GetContacts())
	if err != nil {
		log.Error("cannot update user",
			zap.Int64("id", req.GetId()),
			zap.Error(err),
		)
		return nil, err
	}

	return &desc.UserContactsUpdateResponse{Id: id}, nil
}
