package array

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

// ForEach applies the consumer to each element of the iterator.
func ForEach[T any](iter types.Iterator[T], consumer fx.Consumer[T]) {
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		consumer(v)
	}
}
