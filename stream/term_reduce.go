package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

// Reduce reduces the elements of the iterator to a single value.
// The identity value is the initial value of the reduction and the
// accumulator function combines an element into the partial result.
func Reduce[T any](iterator Iterator[T], identity T, accumulator fx.BinaryOperator[T]) T {
	for {
		v, ok := iterator.Next()
		if !ok {
			break
		}
		identity = accumulator(identity, v)
	}
	return identity
}

// ReduceWithoutIdentity reduces the elements of the iterator to a single value.
// The accumulator function combines an element into the partial result.
// If the iterator is empty, returns false.
func ReduceWithoutIdentity[T any](iterator Iterator[T], accumulator fx.BinaryOperator[T]) (T, bool) {
	v, ok := iterator.Next()
	if !ok {
		return types.Empty[T](), false
	}
	return Reduce[T](iterator, v, accumulator), true
}
