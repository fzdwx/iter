package iter

import (
	"github.com/fzdwx/iter/stream"
)

// Empty returns an empty stream of type T.
//
// Empty does not take any parameters.
// The return type is a pointer to a stream of type T.
func Empty[T any]() *stream.Stream[T] {
	return stream.Empty[T]()
}

// Stream creates a new stream from a given array.
//
// It takes in an array of any type and returns a pointer to a stream.Stream
// object that represents the stream of elements in the array.
// The type of elements in the array can be any type.
func Stream[T any](arr []T) *stream.Stream[T] {
	return stream.Of[T](arr)
}

// Concat takes multiple slices of any type and concatenates them into a single stream.
func Concat[T any](arr ...[]T) *stream.Stream[T] {
	// If there are no input slices, return an empty stream.
	if len(arr) == 0 {
		return Empty[T]()
	}

	// Create an empty stream to store the concatenated values.
	empty := Empty[T]()

	// Iterate over each input slice and concatenate it to the empty stream.
	for _, v := range arr {
		empty = empty.Concat(v)
	}

	// Return the concatenated stream.
	return empty
}
