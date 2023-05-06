package array

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

type commonArrayOps[T any] struct {
	iter types.Iterator[T]
}

func (a *commonArrayOps[T]) Filter(filter fx.Predicate[T]) *filterArray[T] {
	return Filter[T](a.iter, filter)
}

func (a *commonArrayOps[T]) ForEach(consumer fx.Consumer[T]) {
	ForEach[T](a.iter, consumer)
}

func (a *commonArrayOps[T]) ToArray() []T {
	return CollectToArray[T](a.iter)
}

func (a *commonArrayOps[T]) Iter() types.Iterator[T] {
	return a.iter
}

func newCommonArrayOps[T any](iterator types.Iterator[T]) commonArrayOps[T] {
	return commonArrayOps[T]{iter: iterator}
}
