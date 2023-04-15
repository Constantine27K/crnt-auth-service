package auth

import (
	secretStorage "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/storage"
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
	"github.com/Constantine27K/crnt-sdk/pkg/authorization"
	"github.com/Constantine27K/crnt-sdk/pkg/token"
)

type Implementation struct {
	desc.UnimplementedAuthServer
	secretStorage secretStorage.SecretStorage
	// userStorage   userStorage.UserStorage
	tokenMaker token.Maker
	authorizer authorization.Authorizer
}

func NewService(
	secretStorage secretStorage.SecretStorage,
	// userStorage userStorage.UserStorage,
	tokenMaker token.Maker,
	authorizer authorization.Authorizer,
) *Implementation {
	return &Implementation{
		secretStorage: secretStorage,
		//userStorage:   userStorage,
		tokenMaker: tokenMaker,
		authorizer: authorizer,
	}
}
