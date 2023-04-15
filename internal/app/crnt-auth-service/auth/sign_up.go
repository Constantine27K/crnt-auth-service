package auth

import (
	"context"

	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
	log "github.com/sirupsen/logrus"
)

func (i *Implementation) SignUp(ctx context.Context, req *desc.SignUpRequest) (*desc.SignUpResponse, error) {
	id, err := i.secretStorage.Add(req.GetSecret())
	if err != nil {
		log.Error("failed to store secret")
		return nil, err
	}

	return &desc.SignUpResponse{Id: id}, nil
}
