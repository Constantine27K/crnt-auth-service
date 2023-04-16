package auth

import (
	"context"

	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) SignUp(ctx context.Context, req *desc.SignUpRequest) (*desc.SignUpResponse, error) {
	err := i.validator.ValidateSecret(req.GetSecret())
	if err != nil {
		log.Error("failed validate secret")
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	id, err := i.secretStorage.Add(req.GetSecret())
	if err != nil {
		log.Error("failed to store secret")
		return nil, err
	}

	return &desc.SignUpResponse{Id: id}, nil
}
