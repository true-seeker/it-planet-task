package validator

import (
	"errors"
	"fmt"
	"it-planet-task/pkg/config"
	"it-planet-task/pkg/errorHandler"
	"it-planet-task/pkg/paginator"
	"net/http"
	"strconv"
	"time"
)

func ValidateAndReturnPagination(from, size string) (*paginator.Pagination, *errorHandler.HttpErr) {
	pagination := &paginator.Pagination{}
	if from != "" {
		from, httpErr := ValidateAndReturnIntField(from, "from")
		if httpErr != nil {
			return nil, httpErr
		}
		if from < 0 {
			return nil, &errorHandler.HttpErr{
				Err:        errors.New("from must be greater or equal to 0"),
				StatusCode: http.StatusBadRequest,
			}
		}
		pagination.From = from
	}
	if size != "" {
		size, httpErr := ValidateAndReturnIntField(size, "size")
		if httpErr != nil {
			return nil, httpErr
		}
		if size <= 0 {
			return nil, &errorHandler.HttpErr{
				Err:        errors.New("size must be greater than 0"),
				StatusCode: http.StatusBadRequest,
			}
		}
		pagination.Size = size
	}
	return pagination, nil
}

func ValidateAndReturnIntField(field, fieldName string) (int, *errorHandler.HttpErr) {
	intField, err := strconv.Atoi(field)
	if err != nil {
		return 0, &errorHandler.HttpErr{
			Err:        errors.New(fmt.Sprintf("%s must be integer", fieldName)),
			StatusCode: http.StatusBadRequest,
		}
	}
	return intField, nil
}

func ValidateAndReturnId(idStr, fieldName string) (int, *errorHandler.HttpErr) {
	id, httpErr := ValidateAndReturnIntField(idStr, fieldName)
	if httpErr != nil {
		return 0, httpErr
	}

	if id <= 0 {
		return 0, &errorHandler.HttpErr{
			Err:        errors.New(fmt.Sprintf("%s must be greater than 0", fieldName)),
			StatusCode: http.StatusBadRequest,
		}
	}
	return id, nil
}

func ValidateAndReturnDateTime(field, fieldName string) (*time.Time, *errorHandler.HttpErr) {
	date, err := time.Parse(config.DateTimeLayout, field)
	if err != nil {
		return nil, &errorHandler.HttpErr{
			Err:        errors.New(fmt.Sprintf("%s must be in ISO-8601 format", fieldName)),
			StatusCode: http.StatusBadRequest,
		}
	}
	return &date, nil
}

func ValidateAndReturnDate(field, fieldName string) (*time.Time, *errorHandler.HttpErr) {
	date, err := time.Parse(config.DateLayout, field)
	if err != nil {
		return nil, &errorHandler.HttpErr{
			Err:        errors.New(fmt.Sprintf("%s must be in ISO-8601 format", fieldName)),
			StatusCode: http.StatusBadRequest,
		}
	}
	return &date, nil
}
