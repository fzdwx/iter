package main

import (
	"fmt"
	"github.com/fzdwx/iter"
	"github.com/fzdwx/iter/fx"
)

func main() {
	ints := []int{1, 2, 3, 4, 5}
	result := iter.Stream(ints).Reduce(0, fx.Add[int])
	fmt.Println(result) // 15

	result = iter.Stream(ints).Reduce(1, fx.Mul[int])
	fmt.Println(result) // 120

	result, _ = iter.Stream(ints).ReduceWithoutIdentity(fx.Add[int])
	fmt.Println(result) // 15

}
