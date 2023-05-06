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

func (a *commonArrayOps[T]) Iter() types.Iterator[T] {
	return a.iter
}

func newCommonArrayOps[T any](iterator types.Iterator[T]) commonArrayOps[T] {
	return commonArrayOps[T]{iter: iterator}
}

func (a *commonArrayOps[T]) GroupByStr(group fx.Func[T, string]) map[string][]T {
	return GroupBy[T, string](a.iter, group)
}

func (a *commonArrayOps[T]) GroupByInt(group fx.Func[T, int]) map[int][]T {
	return GroupBy[T, int](a.iter, group)
}

func (a *commonArrayOps[T]) GroupByInt8(group fx.Func[T, int8]) map[int8][]T {
	return GroupBy[T, int8](a.iter, group)
}

func (a *commonArrayOps[T]) GroupByInt16(group fx.Func[T, int16]) map[int16][]T {
	return GroupBy[T, int16](a.iter, group)
}

func (a *commonArrayOps[T]) GroupByInt32(group fx.Func[T, int32]) map[int32][]T {
	return GroupBy[T, int32](a.iter, group)
}

func (a *commonArrayOps[T]) GroupByInt64(group fx.Func[T, int64]) map[int64][]T {
	return GroupBy[T, int64](a.iter, group)
}
