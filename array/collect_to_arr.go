package array

import "github.com/fzdwx/iter/types"

func CollectToArray[T any](iter types.Iterator[T]) []T {
	var arr []T
	for {
		v, ok := iter.Next()
		if !ok {
			return arr
		}
		arr = append(arr, v)
	}
}
