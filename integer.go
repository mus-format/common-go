package common

// These constants define the number of bytes required for all uint, int,
// float types in Raw encoding.
const (
	Num64RawSize = 8
	Num32RawSize = 4
	Num16RawSize = 2
	Num8RawSize  = 1
)

// Integer64 is a constraint that permits any 64-bit integer type.
type Integer64 interface {
	~uint | ~uint64 | ~int | ~int64
}

// Integer32 is a constraint that permits any 32-bit integer type.
type Integer32 interface {
	~uint | ~uint32 | ~int | ~int32
}

// Integer16 is a constraint that permits any 16-bit integer type.
type Integer16 interface {
	~uint16 | ~int16
}

// Integer8 is a constraint that permits any 8-bit integer type.
type Integer8 interface {
	~uint8 | ~int8
}

// Num64 is a constraint that permits any 64-bit integer or float type.
type Num64 interface {
	~float64 | Integer64
}

// Num32 is a constraint that permits any 32-bit integer or float type.
type Num32 interface {
	~float32 | Integer32
}
