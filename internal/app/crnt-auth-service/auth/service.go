package auth

import desc "github.com/Constantine27K/crnt-auth-service/pkg/api/auth"

type Implementation struct {
	desc.UnimplementedAuthServer
}

func NewService() *Implementation {
	return &Implementation{}
}
