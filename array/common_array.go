package array

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

type commonArrayOps[T any] struct {
	iter types.Iterator[T]
}

func (a *commonArrayOps[T]) Filter(filter fx.Predicate[T]) *filterArray[T] {
	return Filter[T](a.iter, filter)
}

func (a *commonArrayOps[T]) ForEach(consumer fx.Consumer[T]) {
	ForEach[T](a.iter, consumer)
}

func (a *commonArrayOps[T]) ToArray() []T {
	return CollectToArray[T](a.iter)
}

func (a *commonArrayOps[T]) Iter() types.Iterator[T] {
	return a.iter
}

func newCommonArrayOps[T any](iterator types.Iterator[T]) commonArrayOps[T] {
	return commonArrayOps[T]{iter: iterator}
}

func (a *commonArrayOps[T]) DistinctInt(keyBy fx.Func[T, int]) *distinctArray[T, int] {
	return Distinct[T, int](a.iter, keyBy)
}
func (a *commonArrayOps[T]) DistinctInt8(keyBy fx.Func[T, int8]) *distinctArray[T, int8] {
	return Distinct[T, int8](a.iter, keyBy)
}
func (a *commonArrayOps[T]) DistinctInt16(keyBy fx.Func[T, int16]) *distinctArray[T, int16] {
	return Distinct[T, int16](a.iter, keyBy)
}
func (a *commonArrayOps[T]) DistinctInt32(keyBy fx.Func[T, int32]) *distinctArray[T, int32] {
	return Distinct[T, int32](a.iter, keyBy)
}
func (a *commonArrayOps[T]) DistinctInt64(keyBy fx.Func[T, int64]) *distinctArray[T, int64] {
	return Distinct[T, int64](a.iter, keyBy)
}
func (a *commonArrayOps[T]) DistinctFloat32(keyBy fx.Func[T, float32]) *distinctArray[T, float32] {
	return Distinct[T, float32](a.iter, keyBy)
}
func (a *commonArrayOps[T]) DistinctFloat64(keyBy fx.Func[T, float64]) *distinctArray[T, float64] {
	return Distinct[T, float64](a.iter, keyBy)
}
func (a *commonArrayOps[T]) DistinctUint(keyBy fx.Func[T, uint]) *distinctArray[T, uint] {
	return Distinct[T, uint](a.iter, keyBy)
}
func (a *commonArrayOps[T]) DistinctUint8(keyBy fx.Func[T, uint8]) *distinctArray[T, uint8] {
	return Distinct[T, uint8](a.iter, keyBy)
}
func (a *commonArrayOps[T]) DistinctUint16(keyBy fx.Func[T, uint16]) *distinctArray[T, uint16] {
	return Distinct[T, uint16](a.iter, keyBy)
}
func (a *commonArrayOps[T]) DistinctUint32(keyBy fx.Func[T, uint32]) *distinctArray[T, uint32] {
	return Distinct[T, uint32](a.iter, keyBy)
}
func (a *commonArrayOps[T]) DistinctUint64(keyBy fx.Func[T, uint64]) *distinctArray[T, uint64] {
	return Distinct[T, uint64](a.iter, keyBy)
}
func (a *commonArrayOps[T]) Distinct(keyBy fx.Func[T, string]) *distinctArray[T, string] {
	return Distinct[T, string](a.iter, keyBy)
}

func (a *commonArrayOps[T]) GroupBy(group fx.Func[T, string]) map[string][]T {
	return GroupBy[T, string](a.iter, group)
}

func (a *commonArrayOps[T]) GroupByInt(group fx.Func[T, int]) map[int][]T {
	return GroupBy[T, int](a.iter, group)
}

func (a *commonArrayOps[T]) GroupByInt8(group fx.Func[T, int8]) map[int8][]T {
	return GroupBy[T, int8](a.iter, group)
}

func (a *commonArrayOps[T]) GroupByInt16(group fx.Func[T, int16]) map[int16][]T {
	return GroupBy[T, int16](a.iter, group)
}

func (a *commonArrayOps[T]) GroupByInt32(group fx.Func[T, int32]) map[int32][]T {
	return GroupBy[T, int32](a.iter, group)
}

func (a *commonArrayOps[T]) GroupByInt64(group fx.Func[T, int64]) map[int64][]T {
	return GroupBy[T, int64](a.iter, group)
}

func (a *commonArrayOps[T]) MapToInt(mapper fx.Func[T, int]) *mapArray[T, int] {
	return Map[T, int](a.iter, mapper)
}
func (a *commonArrayOps[T]) MapToInt8(mapper fx.Func[T, int8]) *mapArray[T, int8] {
	return Map[T, int8](a.iter, mapper)
}
func (a *commonArrayOps[T]) MapToInt16(mapper fx.Func[T, int16]) *mapArray[T, int16] {
	return Map[T, int16](a.iter, mapper)
}
func (a *commonArrayOps[T]) MapToInt32(mapper fx.Func[T, int32]) *mapArray[T, int32] {
	return Map[T, int32](a.iter, mapper)
}
func (a *commonArrayOps[T]) MapToInt64(mapper fx.Func[T, int64]) *mapArray[T, int64] {
	return Map[T, int64](a.iter, mapper)
}
func (a *commonArrayOps[T]) MapToFloat32(mapper fx.Func[T, float32]) *mapArray[T, float32] {
	return Map[T, float32](a.iter, mapper)
}
func (a *commonArrayOps[T]) MapToFloat64(mapper fx.Func[T, float64]) *mapArray[T, float64] {
	return Map[T, float64](a.iter, mapper)
}
func (a *commonArrayOps[T]) MapToUint(mapper fx.Func[T, uint]) *mapArray[T, uint] {
	return Map[T, uint](a.iter, mapper)
}
func (a *commonArrayOps[T]) MapToUint8(mapper fx.Func[T, uint8]) *mapArray[T, uint8] {
	return Map[T, uint8](a.iter, mapper)
}
func (a *commonArrayOps[T]) MapToUint16(mapper fx.Func[T, uint16]) *mapArray[T, uint16] {
	return Map[T, uint16](a.iter, mapper)
}
func (a *commonArrayOps[T]) MapToUint32(mapper fx.Func[T, uint32]) *mapArray[T, uint32] {
	return Map[T, uint32](a.iter, mapper)
}
func (a *commonArrayOps[T]) MapToUint64(mapper fx.Func[T, uint64]) *mapArray[T, uint64] {
	return Map[T, uint64](a.iter, mapper)
}
func (a *commonArrayOps[T]) MapTo(mapper fx.Func[T, string]) *mapArray[T, string] {
	return Map[T, string](a.iter, mapper)
}
