package main

import (
	"fmt"
	"github.com/fzdwx/iter"
	"github.com/fzdwx/iter/array"
	"github.com/fzdwx/iter/fx"
)

func main() {
	ints := []int{1, 1, 2, 2, 3, 6, 7, 8, 9, 10}

	a := iter.Array(ints).
		DistinctInt(fx.IdentityInt).
		Filter(func(i int) bool {
			return i > 2
		}).
		MapToInt(func(i int) int {
			return i * 2
		})

	a.ForEach(fx.Println[int])

	m1 := array.ToMap[int, string](iter.Array(ints), func(i int) string {
		return fmt.Sprintf("%d", i)
	})

	m2 := array.ToMapWithValue[int, string, string](iter.Array(ints), func(i int) string {
		return fmt.Sprintf("%d", i)
	}, func(i int) string {
		return fmt.Sprintf("%d", i*2)
	})

	fmt.Println(m1)
	fmt.Println(m2)
}
