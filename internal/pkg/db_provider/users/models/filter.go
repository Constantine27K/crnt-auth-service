package models

import (
	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/user"
	sq "github.com/Masterminds/squirrel"
)

type UsersFilter struct {
	IDs          []int64
	Names        []string
	LastNames    []string
	DisplayNames []string
	Emails       []string
}

func NewUsersFilter(req *desc.UserGetRequest) *UsersFilter {
	return &UsersFilter{
		IDs:          req.GetIds(),
		Names:        req.GetNames(),
		LastNames:    req.GetLastNames(),
		DisplayNames: req.GetDisplayNames(),
		Emails:       req.GetEmails(),
	}
}

func (f *UsersFilter) Apply(query sq.SelectBuilder) sq.SelectBuilder {
	if len(f.IDs) > 0 {
		query = query.Where(sq.Eq{
			"id": f.IDs,
		})
	}

	if len(f.Names) > 0 {
		query = query.Where(sq.Eq{
			"name": f.Names,
		})
	}

	if len(f.LastNames) > 0 {
		query = query.Where(sq.Eq{
			"last_name": f.LastNames,
		})
	}

	if len(f.DisplayNames) > 0 {
		query = query.Where(sq.Eq{
			"display_name": f.DisplayNames,
		})
	}

	if len(f.Emails) > 0 {
		query = query.Where(sq.Eq{
			"email": f.Emails,
		})
	}

	return query
}
