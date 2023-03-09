package validation

import (
	"fmt"
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
)

type Validator interface {
	ValidateSecret(secret *desc.Secret) error
}

type validator struct {
	secret SecretValidator
}

func NewValidator() Validator {
	return &validator{
		secret: NewSecretValidator(),
	}
}

func (v *validator) ValidateSecret(secret *desc.Secret) error {
	if secret == nil {
		return fmt.Errorf("no secret provided")
	}

	return v.secret.Validate(secret.GetLogin(), secret.GetPassword(), secret.GetRole())
}
