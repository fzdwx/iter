package array

import "github.com/fzdwx/iter/types"

type collectToArray[T any] struct {
	iter types.Iterator[T]
}

func newCollectToArray[T any](iter types.Iterator[T]) *collectToArray[T] {
	return &collectToArray[T]{iter: iter}
}

func (c *collectToArray[T]) Collect() []T {
	var arr []T
	for {
		v, ok := c.iter.Next()
		if !ok {
			return arr
		}
		arr = append(arr, v)
	}
}
