package paginator

import (
	"gorm.io/gorm"
	"net/url"
	"strconv"
)

func Paginate(q url.Values) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		from, err := strconv.Atoi(q.Get("from"))
		if err != nil {
			from = 0
		}

		size, err := strconv.Atoi(q.Get("size"))
		if err != nil || size <= 0 {
			size = 10
		}

		return db.Offset(from).Limit(size)
	}
}
