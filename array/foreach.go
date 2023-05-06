package array

import "github.com/fzdwx/iter/types"

func ForEach[T any](iter types.Iterator[T], fn func(T)) {
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		fn(v)
	}
}
