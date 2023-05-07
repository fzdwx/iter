package fx

import (
	"fmt"
	"github.com/fzdwx/iter/types"
)

func Println[T any](t T) {
	fmt.Println(t)
}

func Add[T types.Number](a, b T) T {
	return a + b
}

func Sub[T types.Number](a, b T) T {
	return a - b
}

func Mul[T types.Number](a, b T) T {
	return a * b
}

func Div[T types.Number](a, b T) T {
	return a / b
}
