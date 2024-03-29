package main

import (
	"fmt"
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/seq"
	"github.com/fzdwx/iter/stream"
)

func main() {
	seq.Generator[int](func(accept fx.Consumer[int]) {
		accept(1)
		accept(2)
		accept(3)
	}).Consume(func(i int) {
		println(i)
	})

	f := func() *seq.Impl[int64] {
		return seq.Generator[int64](func(accept fx.Consumer[int64]) {
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
	}

	join := f().Take(10).Join(",", func(i int64) string {
		return fmt.Sprintf("%d", i)
	})

	seq.Map[int64, string](f().Take(10), func(i int64) string {
		return fmt.Sprintf("%d", i)
	}).Consume(func(s string) {
		fmt.Println(s)
	})

	fmt.Println(join)

	iter := stream.Of([]string{"a", "b", "c"}).Iter()
	seq.Zip(seq.Slice([]int{1, 2, 3}).Seq(), iter, func(a int, b string) string {
		return fmt.Sprintf("%d%s", a, b)
	}).Consume(fx.Println[string])
	// 输出：[1a 2b 3c]

	collect := f().Take(10).Collect()
	fmt.Println(collect)
}
