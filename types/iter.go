package types

type Iterator[T any] interface {
	Next() (T, bool)
}
