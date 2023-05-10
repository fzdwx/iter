package main

import (
	"fmt"
	"github.com/fzdwx/iter"
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/stream"
)

func main() {
	ints := []int{1, 1, 2, 2, 3, 6, 7, 8, 9, 10}

	a := iter.Stream(ints).
		DistinctInt(fx.IdentityInt).
		Filter(func(i int) bool {
			return i > 2
		}).
		MapToInt(func(i int) int {
			return i * 2
		})

	a.ForEach(fx.Println[int])

	m1 := stream.ToMap[int, string, string](iter.Stream(ints), func(i int) string {
		return fmt.Sprintf("%d", i)
	}, func(i int) string {
		return fmt.Sprintf("%d", i*2)
	})

	m2 := stream.ToMap2[int, string](iter.Stream(ints), func(i int) string {
		return fmt.Sprintf("%d", i)
	})

	fmt.Println(m1)
	fmt.Println(m2)
}
