package array

import "github.com/fzdwx/iter/types"

type distinctArray[T any, K comparable] struct {
	iter types.Iterator[T]
	m    map[K]bool
	commonArrayOps[T]
}
