package errors

import "errors"

var (
	ErrParseSize       = errors.New("Size is not integer")
	ErrSizeOfSize      = errors.New("Size is too big")
	ErrParseField      = errors.New("Field position number is not integer 0 to 9")
	ErrParseStart      = errors.New("Start position is not integer")
	ErrStartOutOfRange = errors.New("Start position is out of range")
	ErrParseEnd        = errors.New("End position is not integer")
	ErrEndOutOfRange   = errors.New("End position is out of range")
)
