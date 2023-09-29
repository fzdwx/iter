package stream

import "github.com/fzdwx/iter/fx"

func (a *commonStreamOps[T]) FlatMapTo(mapper fx.Func[T, Iterator[string]]) *flatMapStream[T, string] {
	return FlatMap[T, string](a.iter, mapper)
}
