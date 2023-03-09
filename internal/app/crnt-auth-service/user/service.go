package user

import (
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/authorization"
	secretstorage "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/storage"
	userstorage "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/users/storage"
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/validation"
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"
)

type Implementation struct {
	userStorage   userstorage.UserStorage
	secretStorage secretstorage.SecretStorage
	validator     validation.Validator
	authorizer    authorization.Authorizer
	desc.UnimplementedUserRegistryServer
}

func NewService(
	storage userstorage.UserStorage,
	secretStorage secretstorage.SecretStorage,
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
