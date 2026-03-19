package test

import (
	"reflect"
)

func ComparePtrs(t, v any) bool {
	p1 := reflect.ValueOf(t).Pointer()
	p2 := reflect.ValueOf(v).Pointer()
	return p1 == p2
}
