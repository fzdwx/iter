package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat(t *testing.T) {
	i1 := []int{1}
	i2 := []int{2}
	i3 := []int{3}

	assert.Equal(t, []int{1, 2, 3}, Concat(i1, i2, i3).ToArray())
}
