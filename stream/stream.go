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

func Empty[T any]() *Stream[T] {
	return Of[T]([]T{})
}

func Iter[T any](arr []T) Iterator[T] {
	return Of[T](arr)
}

func (a *Stream[T]) Next() (T, bool) {
	a.idx++
	if a.HasNext() {
		return a.arr[a.idx], true
	}
	return types.Empty[T](), false
}

func (a *Stream[T]) skip(n int64) Iterator[T] {
	a.idx += int(n)
	return a.Iter()
}

func (a *Stream[T]) HasNext() bool {
	return a.idx < len(a.arr)
}

// Concat concatenates the given arrays to the end of this array.
// if called after Next(), it will return itself.
func (a *Stream[T]) Concat(others ...[]T) *Stream[T] {
	if a.callNext() {
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

func (a *Stream[T]) UnOrderRemove(idx int64) *Stream[T] {
	if a.callNext() {
		return a
	}
	if idx < 0 || idx >= int64(len(a.arr)) {
		return a
	}

	a.arr[idx] = a.arr[len(a.arr)-1]
	a.arr = a.arr[:len(a.arr)-1]
	return a
}

func (a *Stream[T]) callNext() bool {
	return a.idx != -1
}

func (a *Stream[T]) ToArray() []T {
	return a.arr
}
