package filter

import (
	"gorm.io/gorm"
	"it-planet-task/internal/app/validator"
	"it-planet-task/internal/app/validator/AnimalValidator"
	"it-planet-task/pkg/errorHandler"
	"it-planet-task/pkg/paginator"
	"net/url"
	"time"
)

// AnimalFilterParams Фильтр поиска по животным
type AnimalFilterParams struct {
	StartDateTime      *time.Time
	EndDateTime        *time.Time
	ChipperId          int
	ChippingLocationId int
	LifeStatus         string
	Gender             string

	Pagination paginator.Pagination
}

func (a *AnimalFilterParams) GetPagination() *paginator.Pagination {
	return &a.Pagination
}

// NewAnimalFilterParams Конструктор фильтра
func NewAnimalFilterParams(q url.Values) (*AnimalFilterParams, *errorHandler.HttpErr) {
	params := &AnimalFilterParams{}
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

	if q.Get("chipperId") != "" {
		chipperId, httpErr := validator.ValidateAndReturnId(q.Get("chipperId"), "chipperId")
		if httpErr != nil {
			return nil, httpErr
		}
		params.ChipperId = chipperId
	}

	if q.Get("chippingLocationId") != "" {
		chipperLocationId, httpErr := validator.ValidateAndReturnId(q.Get("chipperLocationId"), "chipperLocationId")
		if httpErr != nil {
			return nil, httpErr
		}
		params.ChippingLocationId = chipperLocationId
	}

	if q.Get("lifeStatus") != "" {
		httpErr := AnimalValidator.ValidateLifeStatus(q.Get("lifeStatus"))
		if httpErr != nil {
			return nil, httpErr
		}
		params.LifeStatus = q.Get("lifeStatus")
	}

	if q.Get("gender") != "" {
		httpErr := AnimalValidator.ValidateGender(q.Get("gender"))
		if httpErr != nil {
			return nil, httpErr
		}
		params.Gender = q.Get("gender")
	}

	pagination, httpErr := validator.ValidateAndReturnPagination(q.Get("from"), q.Get("size"))
	if httpErr != nil {
		return nil, httpErr
	}

	params.Pagination = *pagination
	return params, nil
}

// AnimalFilter Фильтрация
func AnimalFilter(a *AnimalFilterParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if a.StartDateTime != nil {
			db = db.Where("chipping_date_time >= ?", a.StartDateTime)
		}

		if a.EndDateTime != nil {
			db = db.Where("chipping_date_time <= ?", a.EndDateTime)
		}

		if a.ChipperId != 0 {
			db = db.Where("chipper_id = ?", a.ChipperId)
		}

		if a.ChippingLocationId != 0 {
			db = db.Where("chipping_location_id = ?", a.ChippingLocationId)
		}

		if a.LifeStatus != "" {
			db = db.Where("life_status = ?", a.LifeStatus)
		}

		if a.Gender != "" {
			db = db.Where("gender = ?", a.Gender)
		}

		return db
	}
}
