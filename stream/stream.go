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
	if a.HasNext() {
		return a.arr[a.idx], true
	}
	return types.Empty[T](), false
}

func (a *Stream[T]) HasNext() bool {
	return a.idx < len(a.arr)
}

func (a *Stream[T]) ToArray() []T {
	return a.arr
}
