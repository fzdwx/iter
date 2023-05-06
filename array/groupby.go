package array

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

func GroupBy[T any, K comparable](iter types.Iterator[T], groupBy fx.Func[T, K]) map[K][]T {
	return (&groupArray[T, K]{iter: iter, groupBy: groupBy}).Collect()
}

type groupArray[T any, K comparable] struct {
	iter    types.Iterator[T]
	groupBy fx.Func[T, K]
}

func (a *groupArray[T, K]) Collect() map[K][]T {
	m := make(map[K][]T)
	for {
		v, ok := a.iter.Next()
		if !ok {
			break
		}
		key := a.groupBy(v)
		m[key] = append(m[key], v)
	}
	return m
}
