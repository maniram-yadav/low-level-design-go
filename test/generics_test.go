package test

import (
	"lld/generics"
	"testing"
)

func TestGenerics(t *testing.T) {
	res := generics.Divide(10, 0)
	if val, ok := res.Unwrap(); ok {
		println("Result : ", val)
	}
	res2 := generics.Divide(10, 0)
	if _, ok := res2.Unwrap(); !ok {
		println("Result : ", res2.Error)
	}
}
