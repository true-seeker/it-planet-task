package paginator

import (
	"gorm.io/gorm"
)

type Pagination struct {
	From int
	Size int
}

type PaginationInterface interface {
	GetPagination() *Pagination
}

// Paginate реализация пагинации
func Paginate(q PaginationInterface) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		from := q.GetPagination().From
		if from < 0 {
			from = 0
		}

		size := q.GetPagination().Size
		if size <= 0 {
			size = 10
		}

		return db.Offset(from).Limit(size)
	}
}
