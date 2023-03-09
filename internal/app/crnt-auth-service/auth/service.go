package auth

import (
	secretstorage "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/storage"
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/token"
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
)

type Implementation struct {
	secretStorage secretstorage.SecretStorage
	tokenMaker    token.Maker
	desc.UnimplementedAuthServer
}

func NewService(
	secretStorage secretstorage.SecretStorage,
	tokenMaker token.Maker,
) *Implementation {
	return &Implementation{
		secretStorage: secretStorage,
		tokenMaker:    tokenMaker,
	}
}
