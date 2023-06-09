package stream

import (
	"github.com/fzdwx/iter/fx"
)

type commonStreamOps[T any] struct {
	iter Iterator[T]
}

func (a *commonStreamOps[T]) Skip(n int64) *commonStreamOps[T] {
	iter := a.iter.skip(n)
	a.iter = iter
	return a
}

func (a *commonStreamOps[T]) Limit(limit int64) *limitStream[T] {
	return Limit[T](a.iter, limit)
}

func (a *commonStreamOps[T]) skip(n int64) Iterator[T] {
	return a.iter.skip(n)
}

func (a *commonStreamOps[T]) Iter() Iterator[T] {
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

func (a *commonStreamOps[T]) ToMap(keyMapper fx.Func[T, string]) map[string]T {
	return ToMap2[T, string](a.iter, keyMapper)
}

func newCommonArrayOps[T any](iterator Iterator[T]) commonStreamOps[T] {
	return commonStreamOps[T]{iter: iterator}
}
