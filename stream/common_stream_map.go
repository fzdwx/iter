package stream

import "github.com/fzdwx/iter/fx"

func (a *commonStreamOps[T]) MapToInt(mapper fx.Func[T, int]) *mapStream[T, int] {
	return Map[T, int](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToInt8(mapper fx.Func[T, int8]) *mapStream[T, int8] {
	return Map[T, int8](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToInt16(mapper fx.Func[T, int16]) *mapStream[T, int16] {
	return Map[T, int16](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToInt32(mapper fx.Func[T, int32]) *mapStream[T, int32] {
	return Map[T, int32](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToInt64(mapper fx.Func[T, int64]) *mapStream[T, int64] {
	return Map[T, int64](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToFloat32(mapper fx.Func[T, float32]) *mapStream[T, float32] {
	return Map[T, float32](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToFloat64(mapper fx.Func[T, float64]) *mapStream[T, float64] {
	return Map[T, float64](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToUint(mapper fx.Func[T, uint]) *mapStream[T, uint] {
	return Map[T, uint](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToUint8(mapper fx.Func[T, uint8]) *mapStream[T, uint8] {
	return Map[T, uint8](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToUint16(mapper fx.Func[T, uint16]) *mapStream[T, uint16] {
	return Map[T, uint16](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToUint32(mapper fx.Func[T, uint32]) *mapStream[T, uint32] {
	return Map[T, uint32](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToUint64(mapper fx.Func[T, uint64]) *mapStream[T, uint64] {
	return Map[T, uint64](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapTo(mapper fx.Func[T, string]) *mapStream[T, string] {
	return Map[T, string](a.iter, mapper)
}
