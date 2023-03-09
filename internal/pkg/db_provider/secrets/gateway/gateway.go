package gateway

import (
	"database/sql"

	"github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/models"
	sq "github.com/Masterminds/squirrel"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type SecretsGateway interface {
	Add(secret *models.SecretsRow) (int64, error)
	GetByID(id int64) (*models.SecretsRow, error)
}

type gateway struct {
	db      *sql.DB
	builder sq.StatementBuilderType
}

func NewSecretsGateway(db *sql.DB) SecretsGateway {
	return &gateway{
		db:      db,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

const (
	table = "secrets"
)

var (
	columns = []string{"id", "login", "password"}
)

func (g *gateway) Add(secret *models.SecretsRow) (int64, error) {
	values := []interface{}{
		secret.Login, secret.Password,
	}

	query, args, err := g.builder.Insert(table).
		Columns(columns[1:]...).
		Values(values...).
		Suffix("returning id").ToSql()
	if err != nil {
		log.Error("Gateway.Add query error",
			zap.Any("secret", secret),
			zap.Error(err),
		)
		return 0, err
	}

	var id int64
	err = g.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		log.Error("Gateway.Add scan error",
			zap.Any("secret", secret),
			zap.Error(err),
		)
		return 0, err
	}

	return id, nil
}

func (g *gateway) GetByID(id int64) (*models.SecretsRow, error) {
	query, args, err := g.builder.Select(columns...).
		From(table).
		Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Error("Gateway.GetByID query error",
			zap.Int64("id", id),
			zap.Error(err),
		)
		return nil, err
	}

	var secretRow models.SecretsRow
	err = g.db.QueryRow(query, args...).Scan(
		&secretRow.ID,
		&secretRow.Login,
		&secretRow.Password,
	)
	if err != nil {
		log.Error("Gateway.GetByID scan error",
			zap.Int64("id", id),
			zap.Error(err),
		)
		return nil, err
	}

	return &secretRow, nil
}
