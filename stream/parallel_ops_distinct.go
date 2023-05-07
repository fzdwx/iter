package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

type distinctParallelStream[T any, K comparable] struct {
	iter     types.ParallelIterator[T]
	m        map[K]bool
	distinct fx.Func[T, K]
	source   chan T
}

func (m *distinctParallelStream[T, K]) Generate() {
	m.iter.Generate()
}

func (m *distinctParallelStream[T, K]) Source() chan T {
	m.source = make(chan T, len(m.iter.Source()))
	go func() {
		defer close(m.source)
		for item := range m.source {
			key := m.distinct(item)
			if !m.m[key] {
				m.m[key] = true
				m.source <- item
			}

		}
	}()
	return m.source
}
