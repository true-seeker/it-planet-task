package filter

import (
	"fmt"
	"gorm.io/gorm"
	"it-planet-task/pkg/paginator"
	"net/url"
	"strings"
)

type AccountFilterParams struct {
	FirstName string
	LastName  string
	Email     string

	Pagination paginator.Pagination
}

func NewAccountFilterParams(q url.Values) *AccountFilterParams {
	params := &AccountFilterParams{}
	if q.Get("firstName") != "" {
		params.FirstName = q.Get("firstName")
	}
	if q.Get("lastName") != "" {
		params.LastName = q.Get("lastName")
	}
	if q.Get("email") != "" {
		params.Email = q.Get("email")
	}

	if q.Get("from") != "" {
		params.Pagination.From = q.Get("from")
	}
	if q.Get("size") != "" {
		params.Pagination.Size = q.Get("size")
	}
	return params
}

func (a *AccountFilterParams) GetPagination() *paginator.Pagination {
	return &a.Pagination
}

func AccountFilter(params *AccountFilterParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if params.FirstName != "" {
			db = db.Where("LOWER(first_name) LIKE ?",
				fmt.Sprintf("%%%s%%", strings.ToLower(params.FirstName)))
		}

		if params.LastName != "" {
			db = db.Where("LOWER(last_name) LIKE ?",
				fmt.Sprintf("%%%s%%", strings.ToLower(params.LastName)))
		}

		if params.Email != "" {
			db = db.Where("LOWER(email) LIKE ?",
				fmt.Sprintf("%%%s%%", strings.ToLower(params.Email)))
		}

		return db
	}
}
