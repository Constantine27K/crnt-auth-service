package gateway

import (
	"database/sql"
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/secrets/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SecretsGateway interface {
	Add(secret *models.SecretsRow) (int64, error)
	GetByID(id int64) (*models.SecretsRow, error)
	GetByLogin(login string) (*models.SecretsRow, error)
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
	columns = []string{"id", "login", "password", "role"}
)

func (g *gateway) Add(secret *models.SecretsRow) (int64, error) {
	values := []interface{}{
		secret.Login, secret.Password, secret.Role,
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
		pqErr, ok := err.(*pq.Error)
		if ok {
			if len(pqErr.Constraint) > 0 {
				return 0, status.Error(codes.InvalidArgument, "login already exists")
			}
		}
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
		&secretRow.Role,
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

func (g *gateway) GetByLogin(login string) (*models.SecretsRow, error) {
	query, args, err := g.builder.Select(columns...).
		From(table).
		Where(sq.Eq{"login": login}).ToSql()
	if err != nil {
		log.Error("Gateway.GetByID query error",
			zap.String("login", login),
			zap.Error(err),
		)
		return nil, err
	}

	var secretRow models.SecretsRow
	err = g.db.QueryRow(query, args...).Scan(
		&secretRow.ID,
		&secretRow.Login,
		&secretRow.Password,
		&secretRow.Role,
	)
	if err != nil {
		log.Error("Gateway.GetByID scan error",
			zap.String("login", login),
			zap.Error(err),
		)
		return nil, err
	}

	return &secretRow, nil
}
