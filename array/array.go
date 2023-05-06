package array

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

type Array[T any] struct {
	arr []T
	idx int
	commonArrayOps[T]
}

func New[T any](arr []T) *Array[T] {
	a := &Array[T]{arr: arr, idx: -1}
	a.commonArrayOps = newCommonArrayOps[T](a)
	return a
}

func (a *Array[T]) Iter() types.Iterator[T] {
	return a
}

func Iter[T any](arr []T) types.Iterator[T] {
	return New[T](arr)
}

func (a *Array[T]) Next() (T, bool) {
	a.idx++
	if a.idx >= len(a.arr) {
		return types.Empty[T](), false
	}
	return a.arr[a.idx], true
}

func (a *Array[T]) Filter(filter fx.Predicate[T]) *filterArray[T] {
	return Filter[T](a, filter)
}

func (a *Array[T]) ForEach(consumer fx.Consumer[T]) {
	ForEach[T](a, consumer)
}

func (a *Array[T]) ToArray() []T {
	return a.arr
}
