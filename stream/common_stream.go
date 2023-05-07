package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

type commonStreamOps[T any] struct {
	iter types.Iterator[T]
}

func (a *commonStreamOps[T]) Filter(filter fx.Predicate[T]) *filterArray[T] {
	return Filter[T](a.iter, filter)
}

func (a *commonStreamOps[T]) ForEach(consumer fx.Consumer[T]) {
	ForEach[T](a.iter, consumer)
}

func (a *commonStreamOps[T]) ToArray() []T {
	return CollectToArray[T](a.iter)
}

func (a *commonStreamOps[T]) Iter() types.Iterator[T] {
	return a.iter
}

func newCommonArrayOps[T any](iterator types.Iterator[T]) commonStreamOps[T] {
	return commonStreamOps[T]{iter: iterator}
}
