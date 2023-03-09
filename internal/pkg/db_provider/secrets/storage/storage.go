package storage

import (
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/gateway"
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/models"
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
	"golang.org/x/crypto/bcrypt"
)

type SecretStorage interface {
	Add(secret *desc.Secret) (int64, error)
	GetByID(id int64) (*desc.Secret, error)
	GetByLogin(login string) (*desc.Secret, error)
}

type storage struct {
	gw gateway.SecretsGateway
}

func NewSecretStorage(gw gateway.SecretsGateway) SecretStorage {
	return &storage{
		gw: gw,
	}
}

func (s *storage) Add(secret *desc.Secret) (int64, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(secret.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	row := &models.SecretsRow{
		Login:    secret.GetLogin(),
		Password: string(password),
		Role:     secret.GetRole(),
	}

	return s.gw.Add(row)
}

func (s *storage) GetByID(id int64) (*desc.Secret, error) {
	row, err := s.gw.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &desc.Secret{
		Login:    row.Login,
		Password: row.Password,
		Role:     row.Role,
	}, nil
}

func (s *storage) GetByLogin(login string) (*desc.Secret, error) {
	row, err := s.gw.GetByLogin(login)
	if err != nil {
		return nil, err
	}

	return &desc.Secret{
		Login:    row.Login,
		Password: row.Password,
		Role:     row.Role,
	}, nil
}
