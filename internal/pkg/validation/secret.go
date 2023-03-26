package validation

import (
	"fmt"
	"unicode"
)

const (
	minLoginLength    = 5
	minPasswordLength = 6
)

var (
	availableRoles = map[string]struct{}{
		"admin":    {},
		"lead":     {},
		"employee": {},
	}
)

type SecretValidator interface {
	Validate(login, password, role string) error
}

type secretValidator struct {
}

func NewSecretValidator() SecretValidator {
	return &secretValidator{}
}

func (v *secretValidator) Validate(login, password, role string) error {
	if len(login) < minLoginLength {
		return fmt.Errorf("login is too short, minimum length is %v", minLoginLength)
	}

	if _, ok := availableRoles[role]; !ok {
		return fmt.Errorf("no such role")
	}

	if len(password) < minPasswordLength {
		return fmt.Errorf("password is too short, minimum length is %v", minPasswordLength)
	}

	if !hasUpper(password) {
		return fmt.Errorf("password should have at least one capital letter")
	}

	if !hasDigit(password) {
		return fmt.Errorf("password should have at least one digit")
	}

	return nil
}

func hasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func hasDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false

}
