package common

import (
	"errors"
)

// ErrorPrefix is a prefix for all errors.
const ErrorPrefix = "mus: "

// ErrOverflow happens on Unmarshal when bytes number limit of the type was
// exceeded.
var ErrOverflow = errors.New(ErrorPrefix + "overflow")

// ErrNegativeLength happens on Unmarshal when some data was encoded with
// length and value, and length is negative.
var ErrNegativeLength = errors.New(ErrorPrefix + "negative length")

// ErrWrongFormat happends on Unmarshal when an incorrect format is encountered.
var ErrWrongFormat = errors.New(ErrorPrefix + "wrong format")

// ErrUnsupportedIntSize happens on init, if system int size is not supported.
var ErrUnsupportedIntSize = errors.New(ErrorPrefix + "unsupported IntSize")

// ErrUnknownDTM happens when there is no such DTM in Registry.
var ErrUnknownDTM = errors.New(ErrorPrefix + "unknown DTM")

// ErrWrongTypeVersion happens when the type version from Registry dosn't hold
// a specific type.
var ErrWrongTypeVersion = errors.New(ErrorPrefix + "wrong TypeVersion")

// ErrTooLargeLength happens when the encoded length is too large.
var ErrTooLargeLength = errors.New(ErrorPrefix + "too large length")
