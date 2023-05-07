package types

type Iterator[T any] interface {
	Next() (T, bool)
}

type ParallelIterator[T any] interface {
	Generate()
	Source() chan T
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}
