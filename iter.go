package iter

import (
	"github.com/fzdwx/iter/array"
)

func Array[T any](arr []T) *array.Array[T] {
	return array.New[T](arr)
}
