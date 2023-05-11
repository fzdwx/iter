package seq

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/stream"
)

// Generator creates a Seq from a function that accepts a Consumer.
func Generator[T any](f func(accept fx.Consumer[T])) *Impl[T] {
	return newimpl[T](&quickSeq[T]{f})
}

// Unit creates a Seq from a single value.
func Unit[T any](v T) *Impl[T] {
	return newimpl[T](&unit[T]{t: v})
}

// Slice creates a Seq from a slice.
func Slice[T any](slice []T) *Impl[T] {
	return Iter[T](stream.Of(slice))
}

// Iter creates a Seq from an Iterator.
func Iter[T any](iter stream.Iterator[T]) *Impl[T] {
	return newimpl[T](&iterSeq[T]{iter: iter})
}
