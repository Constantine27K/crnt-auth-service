package user

import (
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/authorization"
	secretStorage "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/storage"
	userStorage "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/users/storage"
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/validation"
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"
)

type Implementation struct {
	desc.UnimplementedUserRegistryServer
	userStorage   userStorage.UserStorage
	secretStorage secretStorage.SecretStorage
	validator     validation.Validator
	authorizer    authorization.Authorizer
}

func NewService(
	storage userStorage.UserStorage,
	secretStorage secretStorage.SecretStorage,
	validator validation.Validator,
	authorizer authorization.Authorizer,
) *Implementation {
	return &Implementation{
		userStorage:   storage,
		secretStorage: secretStorage,
		validator:     validator,
		authorizer:    authorizer,
	}
}
