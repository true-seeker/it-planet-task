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

type AnimalTypeFilterParams struct {
	Type string

	Pagination paginator.Pagination
}

func (a *AnimalTypeFilterParams) GetPagination() *paginator.Pagination {
	return &a.Pagination
}

func NewAnimalTypeFilterParams(q url.Values) (*AnimalTypeFilterParams, *errorHandler.HttpErr) {
	params := &AnimalTypeFilterParams{}
	if q.Get("type") != "" {
		params.Type = q.Get("type")
	}

	pagination, httpErr := validator.ValidateAndReturnPagination(q.Get("from"), q.Get("size"))
	if httpErr != nil {
		return nil, httpErr
	}

	params.Pagination = *pagination
	return params, nil
}

func AnimalTypeFilter(a *AnimalTypeFilterParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if a.Type != "" {
			db = db.Where("LOWER(type) LIKE ?",
				fmt.Sprintf("%%%s%%", strings.ToLower(a.Type)))
		}

		return db
	}
}
