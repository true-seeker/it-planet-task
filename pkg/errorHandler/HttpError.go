package errorHandler

import "fmt"

// HttpErr структура для формирования сообщения об ошибке
type HttpErr struct {
	Err        error
	StatusCode int
}

// todo constructor
func (r *HttpErr) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}
