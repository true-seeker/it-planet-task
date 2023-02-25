package filter

import (
	"gorm.io/gorm"
	"it-planet-task/pkg/config"
	"it-planet-task/pkg/paginator"
	"net/url"
	"time"
)

type AnimalFilterParams struct {
	StartDateTime      string
	EndDateTime        string
	ChipperId          string
	ChippingLocationId string
	LifeStatus         string
	Gender             string

	Pagination paginator.Pagination
}

func (a *AnimalFilterParams) GetPagination() *paginator.Pagination {
	return &a.Pagination
}

func NewAnimalFilterParams(q url.Values) *AnimalFilterParams {
	params := &AnimalFilterParams{}
	if q.Get("startDateTime") != "" {
		params.StartDateTime = q.Get("startDateTime")
	}
	if q.Get("endDateTime") != "" {
		params.EndDateTime = q.Get("endDateTime")
	}
	if q.Get("chipperId") != "" {
		params.ChipperId = q.Get("chipperId")
	}
	if q.Get("chippingLocationId") != "" {
		params.ChippingLocationId = q.Get("chippingLocationId")
	}
	if q.Get("lifeStatus") != "" {
		params.LifeStatus = q.Get("lifeStatus")
	}
	if q.Get("gender") != "" {
		params.Gender = q.Get("gender")
	}

	if q.Get("from") != "" {
		params.Pagination.From = q.Get("from")
	}
	if q.Get("size") != "" {
		params.Pagination.Size = q.Get("size")
	}
	return params
}

func AnimalFilter(a *AnimalFilterParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if a.StartDateTime != "" {
			date, _ := time.Parse(config.TimeLayout, a.StartDateTime)
			db = db.Where("chipping_date_time >= ?", date)
		}

		if a.EndDateTime != "" {
			date, _ := time.Parse(config.TimeLayout, a.EndDateTime)
			db = db.Where("chipping_date_time <= ?", date)
		}

		if a.ChipperId != "" {
			db = db.Where("chipper_id = ?", a.ChipperId)
		}

		if a.ChippingLocationId != "" {
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
