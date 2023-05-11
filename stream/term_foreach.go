package stream

import (
	"github.com/fzdwx/iter/fx"
)

// ForEach applies the consumer to each element of the iterator.
func ForEach[T any](iter Iterator[T], consumer fx.Consumer[T]) {
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		consumer(v)
	}
}
