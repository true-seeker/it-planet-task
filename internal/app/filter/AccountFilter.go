package filter

import (
	"fmt"
	"gorm.io/gorm"
	"it-planet-task/internal/app/validator"
	"it-planet-task/pkg/errorHandler"
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

func NewAccountFilterParams(q url.Values) (*AccountFilterParams, *errorHandler.HttpErr) {
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
	pagination, httpErr := validator.ValidateAndReturnPagination(q.Get("from"), q.Get("size"))
	if httpErr != nil {
		return nil, httpErr
	}

	params.Pagination = *pagination

	return params, nil
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
