package stream

import "github.com/fzdwx/iter/types"

func ParallelCollectToArray[T any](iter types.ParallelIterator[T]) []T {
	iter.Generate()

	var arr []T
	for item := range iter.Source() {
		arr = append(arr, item)
	}
	return arr
}
