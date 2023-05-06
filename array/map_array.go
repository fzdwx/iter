package array

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

type mapArray[T, U any] struct {
	iter   types.Iterator[T]
	mapper fx.Func[T, U]
}

func Map[T, U any](iter types.Iterator[T], mapper fx.Func[T, U]) *mapArray[T, U] {
	return &mapArray[T, U]{
		iter:   iter,
		mapper: mapper,
	}
}

func (m *mapArray[T, U]) Next() (U, bool) {
	v, ok := m.iter.Next()
	if !ok {
		return types.Empty[U](), false
	}
	return m.mapper(v), true
}

func (m *mapArray[T, U]) Filter(filter fx.Predicate[U]) *filterArray[U] {
	return Filter[U](m, filter)
}

func (m *mapArray[T, U]) ForEach(consumer fx.Consumer[U]) {
	ForEach[U](m, consumer)
}

func (m *mapArray[T, U]) ToArray() []U {
	return collectToArray[U](m)
}
