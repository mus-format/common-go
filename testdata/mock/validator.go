package mock

import (
	"reflect"

	"github.com/ymz-ncnk/mok"
)

func NewValidator[T any]() Validator[T] {
	return Validator[T]{mok.New("Validator")}
}

type Validator[T any] struct {
	*mok.Mock
}

func (v Validator[T]) RegisterValidate(fn func(t T) (err error)) Validator[T] {
	v.Register("Validate", fn)
	return v
}

func (v Validator[T]) RegisterNValidate(n int,
	fn func(t T) (err error)) Validator[T] {
	v.RegisterN("Validate", n, fn)
	return v
}

func (v Validator[T]) Validate(t T) (err error) {
	var tVal reflect.Value
	if v := reflect.ValueOf(t); (v.Kind() == reflect.Ptr ||
		v.Kind() == reflect.Interface) && v.IsNil() {
		tVal = reflect.Zero(reflect.TypeOf((*T)(nil)).Elem())
	} else {
		tVal = reflect.ValueOf(t)
	}
	vals, err := v.Call("Validate", tVal)
	if err != nil {
		panic(err)
	}
	err, _ = vals[0].(error)
	return
}
