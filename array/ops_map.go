package array

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

func Map[T, U any](iter types.Iterator[T], mapper fx.Func[T, U]) *mapArray[T, U] {
	m := &mapArray[T, U]{
		iter:   iter,
		mapper: mapper,
	}
	m.commonArrayOps = newCommonArrayOps[U](m)
	return m
}

type mapArray[T, U any] struct {
	iter   types.Iterator[T]
	mapper fx.Func[T, U]
	commonArrayOps[U]
}

func (m *mapArray[T, U]) Next() (U, bool) {
	v, ok := m.iter.Next()
	if !ok {
		return types.Empty[U](), false
	}
	return m.mapper(v), true
}

func (m *mapArray[T, U]) ToArray() []U {
	return CollectToArray[U](m)
}
