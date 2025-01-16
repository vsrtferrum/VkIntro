package errors

import "errors"

var (
	ErrParseSize  = errors.New("Size if not unsigned integer")
	ErrParseField = errors.New("Size if not unsigned integer")
)
