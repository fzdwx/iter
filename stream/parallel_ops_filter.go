package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

type filterParallelStream[T any] struct {
	iter   types.ParallelIterator[T]
	filter fx.Predicate[T]
}

func (f *filterParallelStream[T]) OnNext(predicate func(T) bool) {
	f.iter.OnNext(predicate)
}

func (f *filterParallelStream[T]) Source() chan T {
	return f.iter.Source()
}
