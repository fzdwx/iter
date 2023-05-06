package types

import "reflect"

/*
inspired by https://github.com/samber/lo/blob/2a744340d5ae20d374a2a3b774b1df965b4b0e12/type_manipulation.go
*/

// Empty returns an empty value.
func Empty[T any]() T {
	var zero T
	return zero
}

// ToPtr returns a pointer copy of value.
func ToPtr[T any](x T) *T {
	return &x
}

// EmptyableToPtr returns a pointer copy of value if it's nonzero.
// Otherwise, returns nil pointer.
func EmptyableToPtr[T any](x T) *T {
	// 🤮
	isZero := reflect.ValueOf(&x).Elem().IsZero()
	if isZero {
		return nil
	}

	return &x
}
