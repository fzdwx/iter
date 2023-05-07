package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

type commonStreamOps[T any] struct {
	iter types.Iterator[T]
}

func (a *commonStreamOps[T]) Iter() types.Iterator[T] {
	return a.iter
}

func (a *commonStreamOps[T]) Filter(filter fx.Predicate[T]) *filterStream[T] {
	return Filter[T](a.iter, filter)
}

func (a *commonStreamOps[T]) ForEach(consumer fx.Consumer[T]) {
	ForEach[T](a.iter, consumer)
}

// Reduce reduces the elements of the iterator to a single value.
// The identity value is the initial value of the reduction and the
// accumulator function combines an element into the partial result.
func (a *commonStreamOps[T]) Reduce(identity T, accumulator fx.BinaryOperator[T]) T {
	return Reduce[T](a.iter, identity, accumulator)
}

// ReduceWithoutIdentity reduces the elements of the iterator to a single value.
// The accumulator function combines an element into the partial result.
// If the iterator is empty, returns false.
func (a *commonStreamOps[T]) ReduceWithoutIdentity(accumulator fx.BinaryOperator[T]) (T, bool) {
	return ReduceWithoutIdentity[T](a.iter, accumulator)
}

func (a *commonStreamOps[T]) ToArray() []T {
	return CollectToArray[T](a.iter)
}

func newCommonArrayOps[T any](iterator types.Iterator[T]) commonStreamOps[T] {
	return commonStreamOps[T]{iter: iterator}
}
