package user

import (
	secretstorage "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/storage"
	userstorage "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/users/storage"
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"
)

type Implementation struct {
	userStorage   userstorage.UserStorage
	secretStorage secretstorage.SecretStorage
	desc.UnimplementedUserRegistryServer
}

func NewService(
	storage userstorage.UserStorage,
	secretStorage secretstorage.SecretStorage,
) *Implementation {
	return &Implementation{
		userStorage:   storage,
		secretStorage: secretStorage,
	}
}
