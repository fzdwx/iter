package stream

import "github.com/fzdwx/iter/fx"

func (a *commonStreamOps[T]) FlatMapToInt(mapper fx.Func[T, Iterator[int]]) *flatMapStream[T, int] {
	return FlatMap[T, int](a.iter, mapper)
}
func (a *commonStreamOps[T]) FlatMapToInt8(mapper fx.Func[T, Iterator[int8]]) *flatMapStream[T, int8] {
	return FlatMap[T, int8](a.iter, mapper)
}
func (a *commonStreamOps[T]) FlatMapToInt16(mapper fx.Func[T, Iterator[int16]]) *flatMapStream[T, int16] {
	return FlatMap[T, int16](a.iter, mapper)
}
func (a *commonStreamOps[T]) FlatMapToInt32(mapper fx.Func[T, Iterator[int32]]) *flatMapStream[T, int32] {
	return FlatMap[T, int32](a.iter, mapper)
}
func (a *commonStreamOps[T]) FlatMapToInt64(mapper fx.Func[T, Iterator[int64]]) *flatMapStream[T, int64] {
	return FlatMap[T, int64](a.iter, mapper)
}
func (a *commonStreamOps[T]) FlatMapToFloat32(mapper fx.Func[T, Iterator[float32]]) *flatMapStream[T, float32] {
	return FlatMap[T, float32](a.iter, mapper)
}
func (a *commonStreamOps[T]) FlatMapToFloat64(mapper fx.Func[T, Iterator[float64]]) *flatMapStream[T, float64] {
	return FlatMap[T, float64](a.iter, mapper)
}
func (a *commonStreamOps[T]) FlatMapToUint(mapper fx.Func[T, Iterator[uint]]) *flatMapStream[T, uint] {
	return FlatMap[T, uint](a.iter, mapper)
}
func (a *commonStreamOps[T]) FlatMapToUint8(mapper fx.Func[T, Iterator[uint8]]) *flatMapStream[T, uint8] {
	return FlatMap[T, uint8](a.iter, mapper)
}
func (a *commonStreamOps[T]) FlatMapToUint16(mapper fx.Func[T, Iterator[uint16]]) *flatMapStream[T, uint16] {
	return FlatMap[T, uint16](a.iter, mapper)
}
func (a *commonStreamOps[T]) FlatMapToUint32(mapper fx.Func[T, Iterator[uint32]]) *flatMapStream[T, uint32] {
	return FlatMap[T, uint32](a.iter, mapper)
}
func (a *commonStreamOps[T]) FlatMapToUint64(mapper fx.Func[T, Iterator[uint64]]) *flatMapStream[T, uint64] {
	return FlatMap[T, uint64](a.iter, mapper)
}
func (a *commonStreamOps[T]) FlatMapTo(mapper fx.Func[T, Iterator[string]]) *flatMapStream[T, string] {
	return FlatMap[T, string](a.iter, mapper)
}
