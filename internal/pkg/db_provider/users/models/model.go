package models

import "time"

type UserRow struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	LastName    string    `db:"last_name"`
	DisplayName string    `db:"display_name"`
	Birthday    time.Time `db:"birthday"`
	EmployedAt  time.Time `db:"employed_at"`
	FiredAt     time.Time `db:"fired_at"`
	AboutInfo   string    `db:"about_info"`
	AvatarUrl   string    `db:"avatar_url"`
	PhoneNumber string    `db:"phone_number"`
	Email       string    `db:"email"`
	TelegramUrl string    `db:"telegram_url"`
	DiscordUrl  string    `db:"discord_url"`
	Salary      float64   `db:"salary"`
	IsPieceWage bool      `db:"is_piece_wage"`
	Team        string    `db:"team"`
}
