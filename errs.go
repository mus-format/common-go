package common

import (
	"errors"
	"fmt"
)

// ErrorPrefix is a prefix for all serializer errors.
const ErrorPrefix = "mus: "

// ErrOverflow happens on Unmarshal when bytes number limit of the type was
// exceeded.
var ErrOverflow = errors.New(ErrorPrefix + "overflow")

// ErrNegativeLength happens on Unmarshal when some data was encoded with
// length and value, and length is negative.
var ErrNegativeLength = errors.New(ErrorPrefix + "negative length")

// ErrWrongFormat happens on Unmarshal when an incorrect format is encountered.
var ErrWrongFormat = errors.New(ErrorPrefix + "wrong format")

// ErrUnsupportedIntSize happens on init, if system int size is not supported.
var ErrUnsupportedIntSize = errors.New(ErrorPrefix + "unsupported IntSize")

// ErrTooLargeLength happens when the encoded length is too large.
var ErrTooLargeLength = errors.New(ErrorPrefix + "too large length")

// UnexpectedDTMError is returned when an unexpected DTM is encountered.
type UnexpectedDTMError struct {
	dtm DTM
}

func NewUnexpectedDTMError(dtm DTM) UnexpectedDTMError {
	return UnexpectedDTMError{dtm: dtm}
}

func (e UnexpectedDTMError) DTM() DTM {
	return e.dtm
}

func (e UnexpectedDTMError) Error() string {
	return fmt.Sprintf(ErrorPrefix+"unexpected DTM %d", e.dtm)
}

// WrongDTMError is returned when a typed serializer tries to unmarshal data
// with wrong DTM.
type WrongDTMError struct {
	expected DTM
	actual   DTM
}

func NewWrongDTMError(expected, actual DTM) WrongDTMError {
	return WrongDTMError{expected: expected, actual: actual}
}

func (e WrongDTMError) Expected() DTM {
	return e.expected
}

func (e WrongDTMError) Actual() DTM {
	return e.actual
}

func (e WrongDTMError) Error() string {
	return fmt.Sprintf(ErrorPrefix+"wrong DTM, expected %d, got %d", e.expected, e.actual)
}
