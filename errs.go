package com

import (
	"errors"
)

// ErrOverflow happens on Unmarshal when bytes number limit of the type was
// exceeded.
var ErrOverflow = errors.New("overflow")

// ErrNegativeLength happens on Unmarshal when some data was encoded with
// length and value, and length is negative.
var ErrNegativeLength = errors.New("negative length")

// ErrWrongFormat happends on Unmarshal when an incorrect format is encountered.
var ErrWrongFormat = errors.New("wrong format")

// ErrUnsupportedIntSize happens on init, if system int size is not supported.
var ErrUnsupportedIntSize = errors.New("unsupported IntSize")

// ErrUnknownDTM happens when there is no such DTM in Registry.
var ErrUnknownDTM = errors.New("unknown DTM")

// ErrWrongTypeVersion happens when the type version from Registry dosn't hold
// a specific type.
var ErrWrongTypeVersion = errors.New("wrong TypeVersion")

// ErrTooLargeLength happens when the encoded length is too large.
var ErrTooLargeLength = errors.New("too large length")
