package stream

import "github.com/fzdwx/iter/fx"

func (a *commonStreamOps[T]) GroupBy(group fx.Func[T, string]) map[string][]T {
	return GroupBy[T, string](a.iter, group)
}
