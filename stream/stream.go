package stream

import (
	"github.com/fzdwx/iter/types"
)

type Stream[T any] struct {
	arr []T
	idx int
	commonStreamOps[T]
}

func Of[T any](arr []T) *Stream[T] {
	a := &Stream[T]{arr: arr, idx: -1}
	a.commonStreamOps = newCommonArrayOps[T](a)
	return a
}

func Iter[T any](arr []T) types.Iterator[T] {
	return Of[T](arr)
}

func (a *Stream[T]) Next() (T, bool) {
	a.idx++
	if a.idx >= len(a.arr) {
		return types.Empty[T](), false
	}
	return a.arr[a.idx], true
}

func (a *Stream[T]) ToArray() []T {
	return a.arr
}
