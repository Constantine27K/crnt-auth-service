package models

type SecretsRow struct {
	ID       int64  `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}
