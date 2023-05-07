package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReduce(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	result := Of(ints).Reduce(0, fx.Add[int])
	assert.Equal(t, result, 15)

	result = Of(ints).Reduce(1, fx.Mul[int])
	assert.Equal(t, result, 120)
}
