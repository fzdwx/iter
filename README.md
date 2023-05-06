# Iter

A experimental iterator library for Golang.

## Installation

```bash
go get github.com/fzdwx/iter
```

## Usage

playground: https://go.dev/play/p/-8fD0vvzZds

```go
package array

import (
	"fmt"
	"github.com/fzdwx/iter/array"
)

func main() {
	arr := array.New([]int{1, 2, 3, 4, 5})

	// just convert to another struct
	m := array.Map[int, string](arr.Iter(), func(v int) string {
		return "hello" + fmt.Sprintf("%d", v)
	}).Filter(func(v string) bool {
		return v == "hello1" || v == "hello5"
	}).Filter(func(s string) bool {
		return s == "hello1"
	})
    
	// foreach call `Next` method
	fmt.Println(m.ToArray())
}
```