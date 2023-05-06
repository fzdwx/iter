package array

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

type filterArray[T any] struct {
	iter   types.Iterator[T]
	filter fx.Predicate[T]
	groupWrapper[T]
}

func Filter[T any](iter types.Iterator[T], filter fx.Predicate[T]) *filterArray[T] {
	f := &filterArray[T]{
		iter:   iter,
		filter: filter,
	}
	f.groupWrapper = groupWrapper[T]{f}
	return f
}

func (f *filterArray[T]) Next() (T, bool) {
	for {
		v, ok := f.iter.Next()
		if !ok {
			return types.Empty[T](), false
		}
		if f.filter(v) {
			return v, true
		}
	}
}

func (f *filterArray[T]) Filter(filter fx.Predicate[T]) *filterArray[T] {
	return Filter[T](f, filter)
}

func (f *filterArray[T]) ForEach(consumer fx.Consumer[T]) {
	ForEach[T](f, consumer)
}

func (f *filterArray[T]) ToArray() []T {
	return CollectToArray[T](f)
}
