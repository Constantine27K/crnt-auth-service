package auth

import (
	"context"

	"github.com/Constantine27K/crnt-auth-service/internal/pkg/token"
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Authorize(ctx context.Context, req *desc.AuthRequest) (*desc.AuthResponse, error) {
	var payload *token.Payload
	var err error

	switch req.GetEntity() {
	case "user":
		payload, err = i.authorizer.AuthorizeUser(ctx)
	case "admin":
		payload, err = i.authorizer.AuthorizeAdmin(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "unknown entity")
	}

	if err != nil {
		return nil, err
	}

	return &desc.AuthResponse{Username: payload.Username, Role: payload.Role, Team: payload.Team}, nil
}
