package crnt_user_service

import (
	"context"
	"fmt"

	user_service "github.com/Constantine27K/crnt-user-service/pkg/api/user"
	"google.golang.org/grpc"
)

type Service interface {
	GetUserByLogin(ctx context.Context, login string) (*user_service.User, error)
}

type service struct {
	client user_service.UserRegistryClient
}

func NewService(conn *grpc.ClientConn) Service {
	return &service{
		client: user_service.NewUserRegistryClient(conn),
	}
}

func (s *service) GetUserByLogin(ctx context.Context, login string) (*user_service.User, error) {
	resp, err := s.client.GetUsers(ctx, &user_service.UserGetRequest{
		DisplayName: login,
	})

	if err != nil {
		return nil, err
	}

	if len(resp.GetUsers()) == 0 {
		return nil, fmt.Errorf("no such user")
	}

	return resp.GetUsers()[0], nil
}
