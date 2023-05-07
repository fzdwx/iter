package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

func Map[T, U any](iter types.Iterator[T], mapper fx.Func[T, U]) *mapStream[T, U] {
	m := &mapStream[T, U]{
		iter:   iter,
		mapper: mapper,
	}
	m.commonStreamOps = newCommonArrayOps[U](m)
	return m
}

type mapStream[T, U any] struct {
	iter   types.Iterator[T]
	mapper fx.Func[T, U]
	commonStreamOps[U]
}

func (m *mapStream[T, U]) Next() (U, bool) {
	v, ok := m.iter.Next()
	if !ok {
		return types.Empty[U](), false
	}
	return m.mapper(v), true
}
