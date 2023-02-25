package paginator

import (
	"gorm.io/gorm"
	"strconv"
)

type Pagination struct {
	From string
	Size string
}

type PaginationInterface interface {
	GetPagination() *Pagination
}

func Paginate(q PaginationInterface) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		from, err := strconv.Atoi(q.GetPagination().From)
		if err != nil {
			from = 0
		}

		size, err := strconv.Atoi(q.GetPagination().Size)
		if err != nil || size <= 0 {
			size = 10
		}

		return db.Offset(from).Limit(size)
	}
}
