package stream

import (
	"fmt"
	"github.com/fzdwx/iter/types"
	"testing"
)

func TestParallel(t *testing.T) {

	ints := []int{1, 2, 3, 4, 5}

	var parallel types.ParallelIterator[int] = Of(ints).Parallel()
	pmap := ParallelMap(parallel, func(i int) int {
		return i * 2
	})

	arr := ParallelCollectToArray(pmap)
	fmt.Println(arr)
}
