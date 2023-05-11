package stream

type Iterator[T any] interface {
	Next() (T, bool)
	skip(n int64) Iterator[T]
}
