package array

import (
	"fmt"
	"testing"
)

// Iterator is an interface that allows you to iterate over a collection.

func Test_Arr(t *testing.T) {
	arr := New([]int{1, 2, 3, 4, 5})
	m := Map[int, string](arr.Iter(), func(v int) string {
		return "hello" + fmt.Sprintf("%d", v)
	}).Filter(func(v string) bool {
		return v == "hello1" || v == "hello5"
	}).Filter(func(s string) bool {
		return s == "hello1"
	})

	fmt.Println(m.ToArray())
}
