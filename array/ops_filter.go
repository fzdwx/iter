package array

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

func Filter[T any](iter types.Iterator[T], filter fx.Predicate[T]) *filterArray[T] {
	f := &filterArray[T]{
		iter:   iter,
		filter: filter,
	}
	f.commonArrayOps = newCommonArrayOps[T](f)
	return f
}

type filterArray[T any] struct {
	iter   types.Iterator[T]
	filter fx.Predicate[T]
	commonArrayOps[T]
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
