package stream

import (
	"github.com/fzdwx/iter/fx"
)

func ToMap[T any, K comparable, V any](
	iterator Iterator[T],
	keyMapper fx.Func[T, K],
	valueMapper fx.Func[T, V],
) map[K]V {
	return ToMap3(iterator, func(t T) (K, V) {
		return keyMapper(t), valueMapper(t)
	})
}

func ToMap2[T any, K comparable](
	iterator Iterator[T],
	keyMapper fx.Func[T, K],
) map[K]T {
	return ToMap(iterator, keyMapper, fx.Identity[T])
}

func ToMap3[T any, K comparable, V any](
	iterator Iterator[T],
	mapper func(T) (K, V),
) map[K]V {
	m := make(map[K]V)
	for {
		val, ok := iterator.Next()
		if !ok {
			break
		}
		k, v := mapper(val)
		m[k] = v
	}
	return m
}
