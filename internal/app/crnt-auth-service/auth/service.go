package auth

import (
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/authorization"
	secretStorage "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/storage"
	userStorage "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/users/storage"
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/token"
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
)

type Implementation struct {
	desc.UnimplementedAuthServer
	secretStorage secretStorage.SecretStorage
	userStorage   userStorage.UserStorage
	tokenMaker    token.Maker
	authorizer    authorization.Authorizer
}

func NewService(
	secretStorage secretStorage.SecretStorage,
	userStorage userStorage.UserStorage,
	tokenMaker token.Maker,
	authorizer authorization.Authorizer,
) *Implementation {
	return &Implementation{
		secretStorage: secretStorage,
		userStorage:   userStorage,
		tokenMaker:    tokenMaker,
		authorizer:    authorizer,
	}
}
