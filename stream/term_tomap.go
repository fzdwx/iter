package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

func ToMap[T any, K comparable, V any](
	iterator types.Iterator[T],
	keyMapper fx.Func[T, K],
	valueMapper fx.Func[T, V],
) map[K]V {
	m := make(map[K]V)
	for {
		v, ok := iterator.Next()
		if !ok {
			break
		}
		m[keyMapper(v)] = valueMapper(v)
	}
	return m
}

func ToMap2[T any, K comparable](
	iterator types.Iterator[T],
	keyMapper fx.Func[T, K],
) map[K]T {
	return ToMap(iterator, keyMapper, fx.Identity[T])
}
