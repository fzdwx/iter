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

// Concat concatenates the given arrays to the end of this array.
// if called after Next(), it will return itself.
func (a *Stream[T]) Concat(others ...[]T) *Stream[T] {
	if a.idx != -1 {
		return a
	}
	if others == nil || len(others) == 0 {
		return a
	}

	for _, v := range others {
		a.arr = append(a.arr, v...)
	}
	return a
}

func (a *Stream[T]) ToArray() []T {
	return a.arr
}
