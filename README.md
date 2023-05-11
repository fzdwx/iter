# Iter

A experimental iterator library for Golang.

## Installation

```bash
go get github.com/fzdwx/iter
```

## Usage

playground: https://go.dev/play/p/8YyNuaK95UH

```go
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

```

### Stream

ops:

- `Distinct` https://github.com/fzdwx/iter/blob/main/stream/ops_distinct.go#L8
- `Map` https://github.com/fzdwx/iter/blob/main/stream/ops_map.go#L8
- `Filter` https://github.com/fzdwx/iter/blob/main/stream/ops_filter.go#L8
- `Skip`
- `Limit` https://github.com/fzdwx/iter/blob/main/stream/ops_limit.go#L5

collect:

- ToArray https://github.com/fzdwx/iter/blob/main/stream/term_collect_to_arr.go#L6
- GroupBy https://github.com/fzdwx/iter/blob/main/stream/term_groupby.go#L9
- ToMap  https://github.com/fzdwx/iter/blob/main/stream/term_tomap.go#L8
- Reduce https://github.com/fzdwx/iter/blob/main/stream/term_reduce.go#11
