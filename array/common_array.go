package array

import "github.com/fzdwx/iter/types"

type commonArrayOps[T any] struct {
	iter types.Iterator[T]
	groupWrapper[T]
}

func newCommonArrayOps[T any](iterator types.Iterator[T]) commonArrayOps[T] {
	return commonArrayOps[T]{iter: iterator, groupWrapper: groupWrapper[T]{iter: iterator}}
}
