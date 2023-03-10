package errorHandler

import "log"

// FailOnError аварийное завершение, при появлении ошибки
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
