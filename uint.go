package com

import "strconv"

func init() {
	switch strconv.IntSize {
	case 64:
		uintMaxVarIntLen = Uint64MaxVarintLen
		uintMaxLastByte = Uint64MaxLastByte
	case 32:
		uintMaxVarIntLen = Uint32MaxVarintLen
		uintMaxLastByte = Uint32MaxLastByte
	case 16:
		uintMaxVarIntLen = Uint16MaxVarintLen
		uintMaxLastByte = Uint16MaxLastByte
	case 8:
		uintMaxVarIntLen = Uint8MaxVarintLen
		uintMaxLastByte = Uint8MaxLastByte
	default:
		panic("unsupported IntSize")
	}
}

// These constants define the maximum number of bytes used by Varint encoding for
// all uint types.
const (
	Uint64MaxVarintLen = 10
	Uint32MaxVarintLen = 5
	Uint16MaxVarintLen = 3
	Uint8MaxVarintLen  = 1
)

// These constants define the maximum last byte in Varint encoding for all uint
// types.
const (
	Uint64MaxLastByte byte = 1
	Uint32MaxLastByte byte = 15
	Uint16MaxLastByte byte = 3
	Uint8MaxLastByte  byte = 255
)

var uintMaxVarIntLen int
var uintMaxLastByte byte

// UintMaxVarintLen returns the maximum number of bytes used by Varint encoding
// for uint type.
func UintMaxVarintLen() int {
	return uintMaxVarIntLen
}

// UintMaxLastByte returns the maximum last byte in Varint encoding for uint
// type.
func UintMaxLastByte() byte {
	return uintMaxLastByte
}
