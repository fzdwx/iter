package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistinct(t *testing.T) {
	ints := []int{1, 1, 2, 2, 3}
	arr := Of(ints).DistinctInt(fx.IdentityInt).ToArray()
	assert.Equal(t, []int{1, 2, 3}, arr)
}

func TestDistinctBy(t *testing.T) {
	type user struct {
		name string
		age  int
	}

	users := []user{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{"d", 4},
		{"e", 5},
		{"Jake", 5},
		{"Alice", 1},
		{"Bob", 2},
		{"a", 3},
	}

	arr := Of(users).Distinct(func(u user) string {
		return u.name
	}).ToArray()

	assert.Equal(t, len(arr), 8)
}
