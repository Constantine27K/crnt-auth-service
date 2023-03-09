package user

import (
	"context"

	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func (i *Implementation) UpdateUser(ctx context.Context, req *desc.UserUpdateRequest) (*desc.UserUpdateResponse, error) {
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
