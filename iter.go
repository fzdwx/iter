package iter

import (
	"github.com/fzdwx/iter/stream"
)

func Stream[T any](arr []T) *stream.Stream[T] {
	return stream.Of[T](arr)
}
