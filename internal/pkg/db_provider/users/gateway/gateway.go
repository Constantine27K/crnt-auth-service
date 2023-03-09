package gateway

import (
	"database/sql"
	"fmt"
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/users/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UsersGateway interface {
	Add(user *models.UserRow, secretID int64) (int64, error)
	Get(filter *models.UsersFilter) ([]*models.UserRow, error)
	GetByID(id int64) (*models.UserRow, error)
	Update(user *models.UserRow) (int64, error)
}

type gateway struct {
	db      *sql.DB
	builder sq.StatementBuilderType
}

func NewUsersGateway(db *sql.DB) UsersGateway {
	return &gateway{
		db:      db,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

const (
	tableUsers    = "users"
	tableContacts = "contacts"
)

var (
	columnsUsers = []string{"id", "name", "last_name", "display_name", "birthday",
		"employed_at", "fired_at", "about_info", "avatar_url", "contacts_id",
		"salary", "is_piece_wage", "secrets_id"}

	columnsContacts = []string{"id", "phone_number", "email", "telegram_url", "discord_url"}
)

func (g *gateway) Add(user *models.UserRow, secretID int64) (int64, error) {
	values := []interface{}{
		user.PhoneNumber, user.Email, user.TelegramUrl, user.DiscordUrl,
	}

	query, args, err := g.builder.Insert(tableContacts).
		Columns(columnsContacts[1:]...).
		Values(values...).
		Suffix("returning id").ToSql()
	if err != nil {
		log.Error("Gateway.Add query error",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, err
	}

	var idContact int64
	err = g.db.QueryRow(query, args...).Scan(&idContact)
	if err != nil {
		log.Error("Gateway.Add scan error",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, err
	}

	values = []interface{}{
		user.Name, user.LastName, user.DisplayName, user.Birthday, user.EmployedAt, user.FiredAt,
		user.AboutInfo, user.AvatarUrl, idContact, user.Salary, user.IsPieceWage, secretID,
	}

	query, args, err = g.builder.Insert(tableUsers).
		Columns(columnsUsers[1:]...).
		Values(values...).
		Suffix("returning id").ToSql()
	if err != nil {
		log.Error("Gateway.Add query error",
			zap.Any("user", user),
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
				return 0, status.Error(codes.InvalidArgument, "display_name already exists")
			}
		}
		log.Error("Gateway.Add scan error",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, err
	}

	return id, nil
}

func (g *gateway) Get(filter *models.UsersFilter) ([]*models.UserRow, error) {
	query := g.builder.Select(columnsUsers[:len(columnsUsers)-1]...).
		From(tableUsers)

	if filter != nil {
		query = filter.Apply(query)
	}

	result := make([]*models.UserRow, 0)
	contacts := make([]int64, 0)

	stmt, args, err := query.ToSql()
	if err != nil {
		log.Error("Gateway.Get query error",
			zap.Any("filter", filter),
			zap.Error(err),
		)
		return nil, err
	}

	contactRows, err := g.db.Query(stmt, args...)
	if err != nil {
		log.Error("Gateway.Get query error",
			zap.Any("filter", filter),
			zap.Error(err),
		)
		return nil, err
	}
	defer contactRows.Close()

	for contactRows.Next() {
		var userRow models.UserRow
		var contact int64
		err = contactRows.Scan(
			&userRow.ID,
			&userRow.Name,
			&userRow.LastName,
			&userRow.DisplayName,
			&userRow.Birthday,
			&userRow.EmployedAt,
			&userRow.FiredAt,
			&userRow.AboutInfo,
			&userRow.AvatarUrl,
			&contact,
			&userRow.Salary,
			&userRow.IsPieceWage,
		)
		if err != nil {
			log.Error("Gateway.Get scan error",
				zap.Any("filter", filter),
				zap.Error(err),
			)
			return nil, err
		}
		result = append(result, &userRow)
		contacts = append(contacts, contact)
	}

	if err = contactRows.Err(); err != nil {
		return nil, err
	}

	query = g.builder.Select(columnsContacts[1:]...).
		From(tableContacts).
		Where(sq.Eq{"id": contacts})

	stmt, args, err = query.ToSql()
	if err != nil {
		log.Error("Gateway.Get query error",
			zap.Any("filter", filter),
			zap.Error(err),
		)
		return nil, err
	}

	contactRows, err = g.db.Query(stmt, args...)
	if err != nil {
		log.Error("Gateway.Get query error",
			zap.Any("filter", filter),
			zap.Error(err),
		)
		return nil, err
	}
	defer contactRows.Close()

	var i int
	for contactRows.Next() {
		err = contactRows.Scan(
			&result[i].PhoneNumber,
			&result[i].Email,
			&result[i].TelegramUrl,
			&result[i].DiscordUrl,
		)
		if err != nil {
			log.Error("Gateway.Get scan error",
				zap.Any("filter", filter),
				zap.Error(err),
			)
			return nil, err
		}
		i++
	}

	if err = contactRows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (g *gateway) GetByID(id int64) (*models.UserRow, error) {
	query := g.builder.Select(columnsUsers[:len(columnsUsers)-1]...).
		From(tableUsers).
		Where(sq.Eq{"id": id})

	stmt, args, err := query.ToSql()
	if err != nil {
		log.Error("Gateway.GetByID query error",
			zap.Any("id", id),
			zap.Error(err),
		)
		return nil, err
	}

	var userRow models.UserRow
	var contact int64
	err = g.db.QueryRow(stmt, args...).Scan(
		&userRow.ID,
		&userRow.Name,
		&userRow.LastName,
		&userRow.DisplayName,
		&userRow.Birthday,
		&userRow.EmployedAt,
		&userRow.FiredAt,
		&userRow.AboutInfo,
		&userRow.AvatarUrl,
		&contact,
		&userRow.Salary,
		&userRow.IsPieceWage,
	)
	if err != nil {
		log.Error("Gateway.GetByID query error",
			zap.Any("id", id),
			zap.Error(err),
		)
		return nil, err
	}

	query = g.builder.Select(columnsContacts[1:]...).
		From(tableContacts).
		Where(sq.Eq{"id": contact})

	stmt, args, err = query.ToSql()
	if err != nil {
		log.Error("Gateway.GetByID query error",
			zap.Any("id", id),
			zap.Error(err),
		)
		return nil, err
	}

	err = g.db.QueryRow(stmt, args...).Scan(
		&userRow.PhoneNumber,
		&userRow.Email,
		&userRow.TelegramUrl,
		&userRow.DiscordUrl,
	)
	if err != nil {
		log.Error("Gateway.GetByID query error",
			zap.Any("id", id),
			zap.Error(err),
		)
		return nil, err
	}

	return &userRow, nil
}

func (g *gateway) Update(user *models.UserRow) (int64, error) {
	query := g.builder.Update(tableUsers).
		Where(sq.Eq{"id": user.ID})

	if len(user.Name) > 0 {
		query = query.Set("name", user.Name)
	}

	if len(user.LastName) > 0 {
		query = query.Set("last_name", user.LastName)
	}

	if len(user.DisplayName) > 0 {
		query = query.Set("display_name", user.DisplayName)
	}

	if user.Birthday.Unix() != 0 {
		query = query.Set("birthday", user.Birthday)
	}

	if user.EmployedAt.Unix() != 0 {
		query = query.Set("employed_at", user.EmployedAt)
	}

	if user.FiredAt.Unix() != 0 {
		query = query.Set("fired_at", user.FiredAt)
	}

	if len(user.AboutInfo) > 0 {
		query = query.Set("about_info", user.AboutInfo)
	}

	if len(user.AvatarUrl) > 0 {
		query = query.Set("avatar_url", user.AvatarUrl)
	}

	if user.Salary != 0 {
		query = query.Set("salary", user.Salary)
	}

	query = query.Set("is_piece_wage", user.IsPieceWage)

	stmt, args, err := query.ToSql()
	if err != nil {
		log.Error("Gateway.Update query error",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, err
	}

	res, err := g.db.Exec(stmt, args...)
	if err != nil {
		log.Error("Gateway.Update exec error",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error("Gateway.Update affected error",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.Update no rows affected",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	querySelect := g.builder.Select("contacts_id").
		From(tableUsers).
		Where(sq.Eq{"id": user.ID})

	stmt, args, err = querySelect.ToSql()
	if err != nil {
		log.Error("Gateway.Update query error",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, err
	}

	rows, err := g.db.Query(stmt, args...)
	if err != nil {
		log.Error("Gateway.Get query error",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, err
	}
	defer rows.Close()

	var idContact int64
	for rows.Next() {
		err = rows.Scan(
			&idContact,
		)
		if err != nil {
			log.Error("Gateway.Get scan error",
				zap.Any("user", user),
				zap.Error(err),
			)
			return 0, err
		}
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}

	query = g.builder.Update(tableContacts).
		Where(sq.Eq{"id": idContact})

	if len(user.PhoneNumber) > 0 {
		query = query.Set("phone_number", user.PhoneNumber)
	}

	if len(user.Email) > 0 {
		query = query.Set("email", user.Email)
	}

	if len(user.TelegramUrl) > 0 {
		query = query.Set("telegram_url", user.TelegramUrl)
	}

	if len(user.DiscordUrl) > 0 {
		query = query.Set("discord_url", user.DiscordUrl)
	}

	stmt, args, err = query.ToSql()
	if err != nil {
		log.Error("Gateway.Update query error",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, err
	}

	res, err = g.db.Exec(stmt, args...)
	if err != nil {
		log.Error("Gateway.Update exec error",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err = res.RowsAffected()
	if err != nil {
		log.Error("Gateway.Update affected error",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.Update no rows affected",
			zap.Any("user", user),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	return user.ID, nil
}
