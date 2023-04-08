package filter

import (
	"it-planet-task/internal/app/validator"
	"it-planet-task/pkg/errorHandler"
	"it-planet-task/pkg/paginator"
	"net/http"
	"net/url"
	"time"
)

type AreaAnalyticsFilterParams struct {
	StartDateTime *time.Time
	EndDateTime   *time.Time
	Pagination    paginator.Pagination
}

func (a *AreaAnalyticsFilterParams) GetPagination() *paginator.Pagination {
	return &a.Pagination
}

func NewAreaAnalyticsFilterParams(q url.Values) (*AreaAnalyticsFilterParams, *errorHandler.HttpErr) {
	params := &AreaAnalyticsFilterParams{}

	if q.Get("startDate") != "" {
		startDateTime, httpErr := validator.ValidateAndReturnDate(q.Get("startDate"), "startDate")
		if httpErr != nil {
			return nil, httpErr
		}
		params.StartDateTime = startDateTime
	}

	if q.Get("endDate") != "" {
		endDateTime, httpErr := validator.ValidateAndReturnDate(q.Get("endDate"), "endDate")
		if httpErr != nil {
			return nil, httpErr
		}
		params.EndDateTime = endDateTime
	}

	if params.EndDateTime.Compare(*params.StartDateTime) == -1 {
		return nil, errorHandler.NewHttpErr("endDate must be lower than startDate", http.StatusBadRequest)
	}

	pagination, httpErr := validator.ValidateAndReturnPagination(q.Get("from"), q.Get("size"))
	if httpErr != nil {
		return nil, httpErr
	}

	params.Pagination = *pagination

	return params, nil
}
