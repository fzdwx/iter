package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

func Filter[T any](iter Iterator[T], filter fx.Predicate[T]) *filterStream[T] {
	f := &filterStream[T]{
		iter:   iter,
		filter: filter,
	}
	f.commonStreamOps = newCommonArrayOps[T](f)
	return f
}

type filterStream[T any] struct {
	iter   Iterator[T]
	filter fx.Predicate[T]
	commonStreamOps[T]
}

func (f *filterStream[T]) Next() (T, bool) {
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
