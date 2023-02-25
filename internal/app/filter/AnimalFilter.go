package filter

import (
	"gorm.io/gorm"
	"it-planet-task/pkg/config"
	"net/url"
	"time"
)

func AnimalFilter(q url.Values) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		startDateTime := q.Get("startDateTime")
		if startDateTime != "" {
			date, _ := time.Parse(config.TimeLayout, startDateTime)
			db = db.Where("chipping_date_time >= ?", date)
		}

		endDateTime := q.Get("endDateTime")
		if endDateTime != "" {
			date, _ := time.Parse(config.TimeLayout, endDateTime)
			db = db.Where("chipping_date_time <= ?", date)
		}

		chipperId := q.Get("chipperId")
		if chipperId != "" {
			db = db.Where("chipper_id = ?", chipperId)
		}

		chippingLocationId := q.Get("chippingLocationId")
		if chippingLocationId != "" {
			db = db.Where("chipping_location_id = ?", chippingLocationId)
		}

		lifeStatus := q.Get("lifeStatus")
		if lifeStatus != "" {
			db = db.Where("life_status = ?", lifeStatus)
		}

		gender := q.Get("gender")
		if gender != "" {
			db = db.Where("gender = ?", gender)
		}

		return db
	}
}
