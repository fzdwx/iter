package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

func ParallelFilter[T any](
	iterator types.ParallelIterator[T],
	filter fx.Predicate[T],
) types.ParallelIterator[T] {
	return &filterParallelStream[T]{
		iter:   iterator,
		filter: filter,
	}
}

type filterParallelStream[T any] struct {
	iter   types.ParallelIterator[T]
	filter fx.Predicate[T]
}

func (f *filterParallelStream[T]) Generate() {
	f.iter.Generate()
}

func (f *filterParallelStream[T]) Source() chan T {
	source := make(chan T, len(f.iter.Source()))
	go func() {
		defer close(source)
		for item := range f.iter.Source() {
			if f.filter(item) {
				source <- item
			}
		}
	}()
	return source
}
