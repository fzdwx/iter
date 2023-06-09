package stream

import "github.com/fzdwx/iter/fx"

func (a *commonStreamOps[T]) GroupBy(group fx.Func[T, string]) map[string][]T {
	return GroupBy[T, string](a.iter, group)
}

func (a *commonStreamOps[T]) GroupByInt(group fx.Func[T, int]) map[int][]T {
	return GroupBy[T, int](a.iter, group)
}

func (a *commonStreamOps[T]) GroupByInt8(group fx.Func[T, int8]) map[int8][]T {
	return GroupBy[T, int8](a.iter, group)
}

func (a *commonStreamOps[T]) GroupByInt16(group fx.Func[T, int16]) map[int16][]T {
	return GroupBy[T, int16](a.iter, group)
}

func (a *commonStreamOps[T]) GroupByInt32(group fx.Func[T, int32]) map[int32][]T {
	return GroupBy[T, int32](a.iter, group)
}

func (a *commonStreamOps[T]) GroupByInt64(group fx.Func[T, int64]) map[int64][]T {
	return GroupBy[T, int64](a.iter, group)
}
