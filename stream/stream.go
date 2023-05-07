package stream

import (
	"github.com/fzdwx/iter/types"
)

type Stream[T any] struct {
	arr []T
	idx int
	commonStreamOps[T]

	isParallel bool // mark
}

func Of[T any](arr []T) *Stream[T] {
	a := &Stream[T]{arr: arr, idx: -1, isParallel: false}
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
// if called after Parallel(), it will panic.
func (a *Stream[T]) Concat(others ...[]T) *Stream[T] {
	if a.isParallel {
		panic("parallel stream can not concat")
	}
	if a.idx != -1 {
		panic("can not concat after Next()")
	}
	if others == nil || len(others) == 0 {
		return a
	}

	for _, v := range others {
		a.arr = append(a.arr, v...)
	}
	return a
}

// Parallel returns a parallel stream.
func (a *Stream[T]) Parallel() *ParallelStream[T] {
	return a.ParallelWithWorkers(defaultWorkers)
}

func (a *Stream[T]) ParallelWithWorkers(workers int) *ParallelStream[T] {
	a.isParallel = true
	return OfParallelWithWorkers(a.arr, workers)
}

func (a *Stream[T]) ToArray() []T {
	return a.arr
}
