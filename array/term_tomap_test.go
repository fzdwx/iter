package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToMap(t *testing.T) {
	ints := []int{1, 1, 2, 2, 3, 6, 7, 8, 9, 10}

	m1 := ToMap[int, string](New(ints), func(i int) string {
		return fmt.Sprintf("%d", i)
	})

	m2 := ToMapWithValue[int, string, string](New(ints), func(i int) string {
		return fmt.Sprintf("%d", i)
	}, func(i int) string {
		return fmt.Sprintf("%d", i*2)
	})

	assert.Equal(t, m1["1"], 1)
	assert.Equal(t, m2["1"], "2")
	assert.Equal(t, m2["10"], "20")
}
