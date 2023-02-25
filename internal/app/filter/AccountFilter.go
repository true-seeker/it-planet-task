package filter

import (
	"fmt"
	"gorm.io/gorm"
	"net/url"
	"strings"
)

func AccountFilter(q url.Values) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		firstName := q.Get("firstName")
		if firstName != "" {
			db = db.Where("LOWER(first_name) LIKE ?",
				fmt.Sprintf("%%%s%%", strings.ToLower(firstName)))
		}

		lastName := q.Get("lastName")
		if lastName != "" {
			db = db.Where("LOWER(last_name) LIKE ?",
				fmt.Sprintf("%%%s%%", strings.ToLower(lastName)))
		}

		email := q.Get("email")
		if email != "" {
			db = db.Where("LOWER(email) LIKE ?",
				fmt.Sprintf("%%%s%%", strings.ToLower(email)))
		}

		return db
	}
}
