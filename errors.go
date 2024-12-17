package calculation

import "errors"

var (
	ErrDivisionByZero          = errors.New("division by zero")
	ErrUnequalNumberOfBrackets = errors.New("unequal number of brackets")
	ErrInvalidExpression       = errors.New("invalid expression")
	ErrSomethingWentWrong      = errors.New("something went wrong")
)
