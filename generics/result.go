package generics

import (
	"errors"
	"fmt"
)

type Result[T any, E error] struct {
	Value T
	Error E
}

func Ok[T any, E error](v T) Result[T, E] {
	return Result[T, E]{Value: v}
}

func Err[T any, E error](err E) Result[T, E] {
	return Result[T, E]{Error: err}
}

func (r Result[T, E]) IsOk() bool {
	return errors.Is(r.Error, nil)
}

func (r Result[T, E]) Unwrap() (T, bool) {
	if r.IsOk() {
		return r.Value, true
	}
	var zero T
	return zero, false
}

func Divide(a, b int) Result[int, error] {
	if b == 0 {
		return Err[int, error](fmt.Errorf("divide by zero"))
	}
	return Ok[int, error](a / b)
}
