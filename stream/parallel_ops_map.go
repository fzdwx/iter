package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

func ParallelMap[T, U any](
	iter types.ParallelIterator[T],
	mapper fx.Func[T, U],
) types.ParallelIterator[U] {
	m := &mapParallelStream[T, U]{
		iter:   iter,
		mapper: mapper,
	}
	return m
}

type mapParallelStream[T, U any] struct {
	iter   types.ParallelIterator[T]
	mapper fx.Func[T, U]

	source    chan U
	predicate func(U) bool
}

func (m *mapParallelStream[T, U]) OnNext(predicate func(U) bool) {
	m.iter.OnNext(func(t T) bool {
		return true
	})
	m.predicate = predicate
}

func (m *mapParallelStream[T, U]) Source() chan U {
	m.source = make(chan U, len(m.iter.Source()))
	go func() {
		defer close(m.source)
		for t := range m.iter.Source() {
			u := m.mapper(t)
			if m.predicate(u) {
				m.source <- u
			}
		}
	}()
	return m.source
}
