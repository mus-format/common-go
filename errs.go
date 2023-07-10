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
