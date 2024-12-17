package calculation

import "errors"

var (
	ErrDivisionByZero          = errors.New("division by zero")
	ErrUnequalNumberOfBrackets = errors.New("unequal number of brackets")
	ErrTooManyOperationSigns   = errors.New("too many operation signs")
	ErrSomethingWentWrong      = errors.New("something went wrong")
)
