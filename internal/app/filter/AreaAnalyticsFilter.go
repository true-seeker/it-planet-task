package filter

import (
	"errors"
	"gorm.io/gorm"
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
		return nil, &errorHandler.HttpErr{
			Err:        errors.New("endDate must be lower than startDate"),
			StatusCode: http.StatusBadRequest,
		}
	}

	pagination, httpErr := validator.ValidateAndReturnPagination(q.Get("from"), q.Get("size"))
	if httpErr != nil {
		return nil, httpErr
	}

	params.Pagination = *pagination

	return params, nil
}

func AreaAnalyticsFilter(a *AnimalLocationFilterParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if a.StartDateTime != nil {
			db = db.Where("date_time_of_visit_location_point >= ?", a.StartDateTime)
		}

		if a.EndDateTime != nil {
			db = db.Where("date_time_of_visit_location_point <= ?", a.EndDateTime)
		}

		return db
	}
}

func AreaAnalyticsLastLocationFilter(a *AnimalLocationFilterParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if a.StartDateTime != nil {
			db = db.Where("date_time_of_visit_location_point < ?", a.StartDateTime)
		}

		return db
	}
}
