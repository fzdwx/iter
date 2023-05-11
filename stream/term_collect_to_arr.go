package stream

// CollectToArray collects all elements from the iterator into a slice.
func CollectToArray[T any](iter Iterator[T]) []T {
	var arr []T
	for {
		v, ok := iter.Next()
		if !ok {
			return arr
		}
		arr = append(arr, v)
	}
}
