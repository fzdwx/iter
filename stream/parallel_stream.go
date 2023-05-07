package stream

const (
	defaultWorkers = 16
	minWorkers     = 1
)

type ParallelStream[T any] struct {
	arr     []T
	workers int
	source  chan T
}

func (p *ParallelStream[T]) OnNext(predicate func(T) bool) {
	p.source = make(chan T, len(p.arr))

	go func() {
		defer close(p.source)
		for _, t := range p.arr {
			if predicate(t) {
				p.source <- t
			}
		}
	}()
}

func (p *ParallelStream[T]) Source() chan T {
	return p.source
}

func OfParallel[T any](arr []T) *ParallelStream[T] {
	return OfParallelWithWorkers(arr, defaultWorkers)
}

func OfParallelWithWorkers[T any](arr []T, workers int) *ParallelStream[T] {
	a := &ParallelStream[T]{arr: arr, workers: workers}

	return a
}
