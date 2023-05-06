# Iter

A experimental iterator library for Golang.

## Installation

```bash
go get github.com/fzdwx/iter
```

## Usage

playground: https://go.dev/play/p/5HkLWOLM-fv

```go
package main

import (
	"github.com/fzdwx/iter"
	"github.com/fzdwx/iter/fx"
)

func main() {
	ints := []int{1, 1, 2, 2, 3, 6, 7, 8, 9, 10}

	a := iter.Array(ints).
		// convert to another type: `distinctArray`
		DistinctInt(fx.IdentityInt).
		// convert to another type: `filterArray`
		Filter(func(i int) bool {
			return i > 2
		}).
		// convert to another type: `mapArray`
		MapToInt(func(i int) int {
			return i * 2
		})

	// call `Next` to get next element
	a.ForEach(fx.Println[int])
}

```

### Array

ops:

- `Distinct` https://github.com/fzdwx/iter/blob/main/array/ops_distinct.go#L8
- `Map` https://github.com/fzdwx/iter/blob/main/array/ops_map.go#L8
- `Filter` https://github.com/fzdwx/iter/blob/main/array/ops_filter.go#L8

collect:

- ToArray https://github.com/fzdwx/iter/blob/main/array/term_collect_to_arr.go#L6
- GroupBy https://github.com/fzdwx/iter/blob/main/array/term_groupby.go#L9
- ToMap  https://github.com/fzdwx/iter/blob/main/array/term_tomap.go#L8