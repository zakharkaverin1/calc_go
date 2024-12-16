package calculation

import "errors"

var (
	ErrInvalidExpression             = errors.New("invalid expression")
	ErrDivisionByZero                = errors.New("division by zero")
	ErrUnequalNumberOfBrackets       = errors.New("unequal number of brackets")
	ErrNonSingleCharacterIdentifiers = errors.New("non-single character identifiers")
	ErrTooManyOperationSigns         = errors.New("too many operation signs")
	ErrSomethingWentWrong            = errors.New("something went wrong")
)
