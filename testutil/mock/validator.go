package mock

import (
	"github.com/ymz-ncnk/mok"
)

type ValidateFn[T any] func(t T) (err error)

func NewValidator[T any]() Validator[T] {
	return Validator[T]{mok.New("Validator")}
}

type Validator[T any] struct {
	*mok.Mock
}

func (v Validator[T]) RegisterValidate(fn ValidateFn[T]) Validator[T] {
	v.Register("Validate", fn)
	return v
}

func (v Validator[T]) RegisterNValidate(n int, fn ValidateFn[T]) Validator[T] {
	v.RegisterN("Validate", n, fn)
	return v
}

func (v Validator[T]) Validate(t T) (err error) {
	vals, err := v.Call("Validate", mok.SafeVal[T](t))
	if err != nil {
		panic(err)
	}
	err, _ = vals[0].(error)
	return
}
