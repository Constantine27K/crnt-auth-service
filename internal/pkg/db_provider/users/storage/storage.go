package storage

import (
	"fmt"
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/users/gateway"
	"github.com/Constantine27K/crnt-auth-service/internal/pkg/db_provider/users/models"
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserStorage interface {
	Add(user *desc.User, secretID int64) (int64, error)
	Get(filter *models.UsersFilter) ([]*desc.User, error)
	GetByID(id int64) (*desc.User, error)
	GetByLogin(login string) (*desc.User, error)
	Update(id int64, user *desc.User) (int64, error)
	UpdateContacts(id int64, contacts *desc.Contacts) (int64, error)
}

type storage struct {
	gw gateway.UsersGateway
}

func NewUserStorage(gw gateway.UsersGateway) UserStorage {
	return &storage{
		gw: gw,
	}
}

func (s *storage) Add(user *desc.User, secretID int64) (int64, error) {
	row := &models.UserRow{
		Name:        user.GetName(),
		LastName:    user.GetLastName(),
		DisplayName: user.GetDisplayName(),
		Birthday:    user.GetBirthday().AsTime(),
		EmployedAt:  user.GetEmployedAt().AsTime(),
		FiredAt:     user.GetFiredAt().AsTime(),
		AboutInfo:   user.GetAboutInfo(),
		AvatarUrl:   user.GetAvatarUrl(),
		Salary:      user.GetSalary(),
		IsPieceWage: user.GetIsPieceWage(),
		Team:        user.GetTeam(),
	}

	contacts := user.GetContacts()
	if contacts != nil {
		row.PhoneNumber = contacts.GetPhoneNumber()
		row.Email = contacts.GetEmail()
		row.TelegramUrl = contacts.GetTelegramUrl()
		row.DiscordUrl = contacts.GetDiscordUrl()
	}

	return s.gw.Add(row, secretID)
}

func (s *storage) Get(filter *models.UsersFilter) ([]*desc.User, error) {
	rows, err := s.gw.Get(filter)
	if err != nil {
		return nil, err
	}

	result := make([]*desc.User, 0, len(rows))

	for _, row := range rows {
		result = append(result, &desc.User{
			Id:          row.ID,
			Name:        row.Name,
			LastName:    row.LastName,
			DisplayName: row.DisplayName,
			Birthday:    timestamppb.New(row.Birthday),
			EmployedAt:  timestamppb.New(row.EmployedAt),
			FiredAt:     timestamppb.New(row.FiredAt),
			AboutInfo:   row.AboutInfo,
			AvatarUrl:   row.AvatarUrl,
			Contacts: &desc.Contacts{
				PhoneNumber: row.PhoneNumber,
				Email:       row.Email,
				TelegramUrl: row.TelegramUrl,
				DiscordUrl:  row.DiscordUrl,
			},
			Salary:      row.Salary,
			IsPieceWage: row.IsPieceWage,
			Team:        row.Team,
		})
	}

	return result, err
}

func (s *storage) GetByID(id int64) (*desc.User, error) {
	row, err := s.gw.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &desc.User{
		Id:          row.ID,
		Name:        row.Name,
		LastName:    row.LastName,
		DisplayName: row.DisplayName,
		Birthday:    timestamppb.New(row.Birthday),
		EmployedAt:  timestamppb.New(row.EmployedAt),
		FiredAt:     timestamppb.New(row.FiredAt),
		AboutInfo:   row.AboutInfo,
		AvatarUrl:   row.AvatarUrl,
		Contacts: &desc.Contacts{
			PhoneNumber: row.PhoneNumber,
			Email:       row.Email,
			TelegramUrl: row.TelegramUrl,
			DiscordUrl:  row.DiscordUrl,
		},
		Salary:      row.Salary,
		IsPieceWage: row.IsPieceWage,
		Team:        row.Team,
	}, nil
}

func (s *storage) GetByLogin(login string) (*desc.User, error) {
	row, err := s.gw.GetByLogin(login)
	if err != nil {
		return nil, err
	}

	return &desc.User{
		Id:          row.ID,
		Name:        row.Name,
		LastName:    row.LastName,
		DisplayName: row.DisplayName,
		Birthday:    timestamppb.New(row.Birthday),
		EmployedAt:  timestamppb.New(row.EmployedAt),
		FiredAt:     timestamppb.New(row.FiredAt),
		AboutInfo:   row.AboutInfo,
		AvatarUrl:   row.AvatarUrl,
		Contacts: &desc.Contacts{
			PhoneNumber: row.PhoneNumber,
			Email:       row.Email,
			TelegramUrl: row.TelegramUrl,
			DiscordUrl:  row.DiscordUrl,
		},
		Salary:      row.Salary,
		IsPieceWage: row.IsPieceWage,
		Team:        row.Team,
	}, nil
}

func (s *storage) Update(id int64, user *desc.User) (int64, error) {
	row := &models.UserRow{
		ID:          id,
		Name:        user.GetName(),
		LastName:    user.GetLastName(),
		DisplayName: user.GetDisplayName(),
		Birthday:    user.GetBirthday().AsTime(),
		EmployedAt:  user.GetEmployedAt().AsTime(),
		FiredAt:     user.GetFiredAt().AsTime(),
		AboutInfo:   user.GetAboutInfo(),
		AvatarUrl:   user.GetAvatarUrl(),
		Salary:      user.GetSalary(),
		IsPieceWage: user.GetIsPieceWage(),
		Team:        user.GetTeam(),
	}

	contacts := user.GetContacts()
	if contacts != nil {
		row.PhoneNumber = contacts.GetPhoneNumber()
		row.Email = contacts.GetEmail()
		row.TelegramUrl = contacts.GetTelegramUrl()
		row.DiscordUrl = contacts.GetDiscordUrl()
	}

	return s.gw.Update(row)
}

func (s *storage) UpdateContacts(id int64, contacts *desc.Contacts) (int64, error) {
	if contacts == nil {
		return 0, fmt.Errorf("empty contacts")
	}

	row := &models.UserRow{
		ID: id,
	}

	row.PhoneNumber = contacts.GetPhoneNumber()
	row.Email = contacts.GetEmail()
	row.TelegramUrl = contacts.GetTelegramUrl()
	row.DiscordUrl = contacts.GetDiscordUrl()

	return s.gw.UpdateContacts(row)
}
