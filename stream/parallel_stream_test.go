package stream

import (
	"fmt"
	"github.com/fzdwx/iter/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParallel(t *testing.T) {

	ints := []int{1, 2, 3, 4, 5}

	var parallel types.ParallelIterator[int] = Of(ints).Parallel()
	pmap := ParallelMap(parallel, func(i int) string {
		return fmt.Sprintf("%d", i*2)
	})

	fp := ParallelFilter(pmap, func(s string) bool {
		fmt.Println(s)
		return s != "6"
	})

	arr := ParallelCollectToArray(fp)
	fmt.Println(arr)
	assert.Equal(t, []string{"2", "4", "8", "10"}, arr)
}
