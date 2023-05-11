package main

import (
	"fmt"
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/seq"
)

func main() {
	seq.Generator[int](func(accept fx.Consumer[int]) {
		accept(1)
		accept(2)
		accept(3)
	}).Consume(func(i int) {
		println(i)
	})

	fib := seq.Generator[int64](func(accept fx.Consumer[int64]) {
		var (
			i int64 = 1
			j int64 = 2
		)
		accept(i)
		accept(j)

		for {
			i, j = j, i+j
			accept(j)
		}
	})

	join := fib.Take(10).Join(",", func(i int64) string {
		return fmt.Sprintf("%d", i)
	})
	fmt.Println(join)
}
