package global

import "errors"

var (
	ErrNoDataFound = errors.New("нет данных")
	UnexpectedErr  = errors.New("непредвиденная ошибка")
)
