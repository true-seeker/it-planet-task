package paginator

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Pagination struct {
	From     int
	Size     int
	OrderBy  string
	OrderDir string
}

type PaginationInterface interface {
	GetPagination() *Pagination
}

// Paginate реализация пагинации и фильтрации
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

		order := "id"
		if q.GetPagination().OrderBy != "" {
			orderDir := q.GetPagination().OrderDir
			if orderDir != "" && (strings.ToLower(orderDir) == "asc" || strings.ToLower(orderDir) == "desc") {
				order = fmt.Sprintf("%s %s", q.GetPagination().OrderBy, orderDir)
			} else {
				order = q.GetPagination().OrderBy
			}
		}
		return db.Offset(from).Limit(size).Order(order)
	}
}
