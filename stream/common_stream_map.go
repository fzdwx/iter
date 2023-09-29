package stream

import "github.com/fzdwx/iter/fx"

func (a *commonStreamOps[T]) MapTo(mapper fx.Func[T, string]) *mapStream[T, string] {
	return Map[T, string](a.iter, mapper)
}
