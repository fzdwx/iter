package stream

import "github.com/fzdwx/iter/fx"

func (a *commonStreamOps[T]) DistinctInt(keyBy fx.Func[T, int]) *distinctStream[T, int] {
	return Distinct[T, int](a.iter, keyBy)
}
func (a *commonStreamOps[T]) DistinctInt8(keyBy fx.Func[T, int8]) *distinctStream[T, int8] {
	return Distinct[T, int8](a.iter, keyBy)
}
func (a *commonStreamOps[T]) DistinctInt16(keyBy fx.Func[T, int16]) *distinctStream[T, int16] {
	return Distinct[T, int16](a.iter, keyBy)
}
func (a *commonStreamOps[T]) DistinctInt32(keyBy fx.Func[T, int32]) *distinctStream[T, int32] {
	return Distinct[T, int32](a.iter, keyBy)
}
func (a *commonStreamOps[T]) DistinctInt64(keyBy fx.Func[T, int64]) *distinctStream[T, int64] {
	return Distinct[T, int64](a.iter, keyBy)
}
func (a *commonStreamOps[T]) DistinctFloat32(keyBy fx.Func[T, float32]) *distinctStream[T, float32] {
	return Distinct[T, float32](a.iter, keyBy)
}
func (a *commonStreamOps[T]) DistinctFloat64(keyBy fx.Func[T, float64]) *distinctStream[T, float64] {
	return Distinct[T, float64](a.iter, keyBy)
}
func (a *commonStreamOps[T]) DistinctUint(keyBy fx.Func[T, uint]) *distinctStream[T, uint] {
	return Distinct[T, uint](a.iter, keyBy)
}
func (a *commonStreamOps[T]) DistinctUint8(keyBy fx.Func[T, uint8]) *distinctStream[T, uint8] {
	return Distinct[T, uint8](a.iter, keyBy)
}
func (a *commonStreamOps[T]) DistinctUint16(keyBy fx.Func[T, uint16]) *distinctStream[T, uint16] {
	return Distinct[T, uint16](a.iter, keyBy)
}
func (a *commonStreamOps[T]) DistinctUint32(keyBy fx.Func[T, uint32]) *distinctStream[T, uint32] {
	return Distinct[T, uint32](a.iter, keyBy)
}
func (a *commonStreamOps[T]) DistinctUint64(keyBy fx.Func[T, uint64]) *distinctStream[T, uint64] {
	return Distinct[T, uint64](a.iter, keyBy)
}
func (a *commonStreamOps[T]) Distinct(keyBy fx.Func[T, string]) *distinctStream[T, string] {
	return Distinct[T, string](a.iter, keyBy)
}
