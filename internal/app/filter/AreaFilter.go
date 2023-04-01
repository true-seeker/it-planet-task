package filter

import (
	"it-planet-task/internal/app/validator"
	"it-planet-task/pkg/errorHandler"
	"it-planet-task/pkg/paginator"
	"net/url"
)

type AreaFilterParams struct {
	Pagination paginator.Pagination
}

func (a *AreaFilterParams) GetPagination() *paginator.Pagination {
	return &a.Pagination
}

func NewAreaFilterParams(q url.Values) (*AreaFilterParams, *errorHandler.HttpErr) {
	params := &AreaFilterParams{}
	pagination, httpErr := validator.ValidateAndReturnPagination(q.Get("from"), q.Get("size"))
	if httpErr != nil {
		return nil, httpErr
	}

	params.Pagination = *pagination

	return params, nil
}
