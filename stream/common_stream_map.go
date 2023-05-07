package stream

import "github.com/fzdwx/iter/fx"

func (a *commonStreamOps[T]) MapToInt(mapper fx.Func[T, int]) *mapArray[T, int] {
	return Map[T, int](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToInt8(mapper fx.Func[T, int8]) *mapArray[T, int8] {
	return Map[T, int8](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToInt16(mapper fx.Func[T, int16]) *mapArray[T, int16] {
	return Map[T, int16](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToInt32(mapper fx.Func[T, int32]) *mapArray[T, int32] {
	return Map[T, int32](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToInt64(mapper fx.Func[T, int64]) *mapArray[T, int64] {
	return Map[T, int64](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToFloat32(mapper fx.Func[T, float32]) *mapArray[T, float32] {
	return Map[T, float32](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToFloat64(mapper fx.Func[T, float64]) *mapArray[T, float64] {
	return Map[T, float64](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToUint(mapper fx.Func[T, uint]) *mapArray[T, uint] {
	return Map[T, uint](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToUint8(mapper fx.Func[T, uint8]) *mapArray[T, uint8] {
	return Map[T, uint8](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToUint16(mapper fx.Func[T, uint16]) *mapArray[T, uint16] {
	return Map[T, uint16](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToUint32(mapper fx.Func[T, uint32]) *mapArray[T, uint32] {
	return Map[T, uint32](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapToUint64(mapper fx.Func[T, uint64]) *mapArray[T, uint64] {
	return Map[T, uint64](a.iter, mapper)
}
func (a *commonStreamOps[T]) MapTo(mapper fx.Func[T, string]) *mapArray[T, string] {
	return Map[T, string](a.iter, mapper)
}
