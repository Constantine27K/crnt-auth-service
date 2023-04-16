package auth

import (
	secretStorage "github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/storage"
	userService "github.com/Constantine27K/crnt-auth-service/internal/pkg/services/crnt-user-service"
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/validation"
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
	"github.com/Constantine27K/crnt-sdk/pkg/authorization"
	"github.com/Constantine27K/crnt-sdk/pkg/token"
)

type Implementation struct {
	desc.UnimplementedAuthServer
	secretStorage secretStorage.SecretStorage
	userService   userService.Service
	tokenMaker    token.Maker
	authorizer    authorization.Authorizer
	validator     validation.Validator
}

func NewService(
	secretStorage secretStorage.SecretStorage,
	userService userService.Service,
	tokenMaker token.Maker,
	authorizer authorization.Authorizer,
	validator validation.Validator,
) *Implementation {
	return &Implementation{
		secretStorage: secretStorage,
		userService:   userService,
		tokenMaker:    tokenMaker,
		authorizer:    authorizer,
		validator:     validator,
	}
}
