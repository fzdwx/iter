package fx

import (
	"fmt"
)

type Predicate[T any] func(T) bool

type Func[T, U any] func(T) U

type BiFunc[T, U, V any] func(T, U) V

type BinaryOperator[T any] BiFunc[T, T, T]

type Consumer[T any] func(T)

func Identity[T any](t T) T {
	return t
}

func IdentityInt(t int) int {
	return t
}

func IdentityString(t string) string {
	return t
}

func IdentityBool(t bool) bool {
	return t
}

func IdentityFloat32(t float32) float32 {
	return t
}

func IdentityFloat64(t float64) float64 {
	return t
}

func IdentityInt8(t int8) int8 {
	return t
}

func IdentityInt16(t int16) int16 {
	return t
}

func IdentityInt32(t int32) int32 {
	return t
}

func IdentityInt64(t int64) int64 {
	return t
}

func IdentityUint(t uint) uint {
	return t
}

func IdentityUint8(t uint8) uint8 {
	return t
}

func IdentityUint16(t uint16) uint16 {
	return t
}

func IdentityUint32(t uint32) uint32 {
	return t
}

func IdentityUint64(t uint64) uint64 {
	return t
}

func IdentityByte(t byte) byte {
	return t
}

func IdentityRune(t rune) rune {
	return t
}

func IntToStr(i int) string {
	return fmt.Sprintf("%d", i)
}

func Int8ToStr(i int8) string {
	return fmt.Sprintf("%d", i)
}

func Int16ToStr(i int16) string {
	return fmt.Sprintf("%d", i)
}

func Int32ToStr(i int32) string {
	return fmt.Sprintf("%d", i)
}

func Int64ToStr(i int64) string {
	return fmt.Sprintf("%d", i)
}

func UintToStr(i uint) string {
	return fmt.Sprintf("%d", i)
}

func Uint8ToStr(i uint8) string {
	return fmt.Sprintf("%d", i)
}

func Uint16ToStr(i uint16) string {
	return fmt.Sprintf("%d", i)
}

func Uint32ToStr(i uint32) string {
	return fmt.Sprintf("%d", i)
}

func Uint64ToStr(i uint64) string {
	return fmt.Sprintf("%d", i)
}

func Float32ToStr(i float32) string {
	return fmt.Sprintf("%f", i)
}

func Float64ToStr(i float64) string {
	return fmt.Sprintf("%f", i)
}

func BoolToStr(i bool) string {
	return fmt.Sprintf("%t", i)
}
