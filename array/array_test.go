package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Arr(t *testing.T) {
	arr := Iter([]int{1, 2, 3, 4, 5})
	m := Map[int, string](arr, func(v int) string {
		return "hello" + fmt.Sprintf("%d", v)
	}).Filter(func(v string) bool {
		return v == "hello1" || v == "hello5"
	}).Filter(func(s string) bool {
		return s == "hello1"
	})

	assert.Equal(t, []string{"hello1"}, m.ToArray())
}

func TestMap(t *testing.T) {
	arr := New([]int{1, 2, 3, 4, 5})
	m := Map[int, string](arr.Iter(), func(v int) string {
		return "hello" + fmt.Sprintf("%d", v)
	}).ToArray()

	assert.Equal(t, []string{"hello1", "hello2", "hello3", "hello4", "hello5"}, m)
}

func TestForEach(t *testing.T) {
	arr := New([]int{1, 2, 3, 4, 5})
	idx := 0
	arr.ForEach(func(v int) {
		assert.Equal(t, v, idx+1)
		idx++
	})

	idx = 0
	Map(arr.Iter(), func(v int) string {
		return "hello" + fmt.Sprintf("%d", v)
	}).ForEach(func(v string) {
		assert.Equal(t, v, "hello"+fmt.Sprintf("%d", idx))
		idx++
	})
}
