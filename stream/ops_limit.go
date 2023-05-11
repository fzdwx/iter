package stream

import "github.com/fzdwx/iter/types"

func Limit[T any](iter Iterator[T], limit int64) *limitStream[T] {
	l := &limitStream[T]{
		iter:  iter,
		limit: limit,
	}
	l.commonStreamOps = newCommonArrayOps[T](l)
	return l
}

type limitStream[T any] struct {
	iter  Iterator[T]
	limit int64
	commonStreamOps[T]
}

func (l *limitStream[T]) Next() (T, bool) {
	if l.limit <= 0 {
		return types.Empty[T](), false
	}
	l.limit--
	return l.iter.Next()
}
