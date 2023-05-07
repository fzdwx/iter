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

	source chan U
}

func (m *mapParallelStream[T, U]) Generate() {
	m.iter.Generate()
}

func (m *mapParallelStream[T, U]) Source() chan U {
	m.source = make(chan U, len(m.iter.Source()))
	go func() {
		defer close(m.source)
		for t := range m.iter.Source() {
			m.source <- m.mapper(t)
		}
	}()
	return m.source
}
