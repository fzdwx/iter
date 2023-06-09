package stream

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
	arr := Of([]int{1, 2, 3, 4, 5})
	m := Map[int, string](arr.Iter(), func(v int) string {
		return "hello" + fmt.Sprintf("%d", v)
	}).ToArray()

	assert.Equal(t, []string{"hello1", "hello2", "hello3", "hello4", "hello5"}, m)
}

func TestForEach(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	arr := Of(ints)
	idx := 0
	arr.ForEach(func(v int) {
		assert.Equal(t, v, idx+1)
		idx++
	})

	idx = 1
	Map(Of(ints).Iter(), func(v int) string {
		return "hello" + fmt.Sprintf("%d", v)
	}).ForEach(func(v string) {
		println(v)
		assert.Equal(t, v, "hello"+fmt.Sprintf("%d", idx))
		idx++
	})
}

func TestGroupBy(t *testing.T) {
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
	arr := Of(users)

	m1 := Map(arr.Iter(), func(u user) user {
		return u
	}).GroupBy(func(u user) string {
		return u.name
	})

	assert.Equal(t, 8, len(m1))
	assert.Equal(t, m1["a"], []user{{"a", 1}, {"a", 3}})
	assert.Equal(t, m1["b"], []user{{"b", 2}})

	m2 := Map(Of(users).Iter(), func(u user) user {
		return u
	}).GroupByInt(func(u user) int {
		return u.age
	})

	assert.Equal(t, 5, len(m2))
	assert.Equal(t, m2[1], []user{{"a", 1}, {"Alice", 1}})
}

func TestGroupBy2(t *testing.T) {
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
	m3 := Map(Of(users).Iter(), func(u user) string {
		return u.name
	}).Iter()

	g := GroupBy(m3, func(u string) string {
		return u
	})

	assert.Equal(t, 8, len(g))
	assert.Equal(t, g["a"], []string{"a", "a"})
	assert.Equal(t, g["b"], []string{"b"})
}

func TestDelete(t *testing.T) {
	ints := []int{1, 2, 3, 44, 5}
	arr := Of(ints)
	afterRemove := arr.UnOrderRemove(2).ToArray()

	assert.Equal(t, []int{1, 2, 5, 44, 5}, ints)
	assert.Equal(t, []int{1, 2, 5, 44}, arr.ToArray())
	assert.Equal(t, []int{1, 2, 5, 44}, afterRemove)
}

func TestSkip(t *testing.T) {
	ints := []int{1, 2, 3, 44, 5}
	s := Of(ints).Skip(1).Skip(1).ToArray()

	assert.Equal(t, []int{3, 44, 5}, s)

	s = Of(ints).Skip(1).Skip(1).Skip(1).ToArray()
	assert.Equal(t, []int{44, 5}, s)

	s = Of(ints).Skip(10).ToArray()
	assert.Equal(t, len(s), 0)
}

func TestLimit(t *testing.T) {
	ints := []int{1, 2, 3, 44, 5}
	s := Of(ints).Limit(1).Limit(1).ToArray()

	assert.Equal(t, []int{1}, s)

	s = Of(ints).Limit(1).Limit(1).Limit(1).ToArray()
	assert.Equal(t, []int{1}, s)

	s = Of(ints).Limit(10).ToArray()
	assert.Equal(t, len(s), 5)

	s = Of(ints).Skip(2).Limit(1).ToArray()
	assert.Equal(t, []int{3}, s)
}

func TestFlatMap(t *testing.T) {
	ints := []int{1, 2, 3}
	array := FlatMap[int, int](Of(ints), func(i int) Iterator[int] {
		return Of([]int{i, i + 1})
	}).ToArray()

	assert.Equal(t, []int{1, 2, 2, 3, 3, 4}, array)
}
