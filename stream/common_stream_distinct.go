package stream

import "github.com/fzdwx/iter/fx"

func (a *commonStreamOps[T]) Distinct(keyBy fx.Func[T, string]) *distinctStream[T, string] {
	return Distinct[T, string](a.iter, keyBy)
}
