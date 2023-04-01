package filter

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/validator"
	"it-planet-task/pkg/errorHandler"
	"it-planet-task/pkg/paginator"
	"net/url"
	"time"
)

type AnimalLocationFilterParams struct {
	StartDateTime *time.Time
	EndDateTime   *time.Time

	Pagination paginator.Pagination
}

func (a *AnimalLocationFilterParams) GetPagination() *paginator.Pagination {
	return &a.Pagination
}

func NewAnimalLocationFilterParams(q url.Values) (*AnimalLocationFilterParams, *errorHandler.HttpErr) {
	params := &AnimalLocationFilterParams{}

	if q.Get("startDateTime") != "" {
		startDateTime, httpErr := validator.ValidateAndReturnDateTime(q.Get("startDateTime"), "startDateTime")
		if httpErr != nil {
			return nil, httpErr
		}
		params.StartDateTime = startDateTime
	}

	if q.Get("endDateTime") != "" {
		endDateTime, httpErr := validator.ValidateAndReturnDateTime(q.Get("endDateTime"), "endDateTime")
		if httpErr != nil {
			return nil, httpErr
		}
		params.EndDateTime = endDateTime
	}

	pagination, httpErr := validator.ValidateAndReturnPagination(q.Get("from"), q.Get("size"))
	if httpErr != nil {
		return nil, httpErr
	}
	params.Pagination = *pagination

	return params, nil
}

func AnimalLocationFilter(a *AnimalLocationFilterParams) func(db *gorm.DB) *gorm.DB {
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
