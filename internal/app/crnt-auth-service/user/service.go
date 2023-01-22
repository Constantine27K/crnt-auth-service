package user

import desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"

type Implementation struct {
	desc.UnimplementedUserRegistryServer
}

func NewService() *Implementation {
	return &Implementation{}
}
