package array

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

func Distinct[T any, K comparable](iter types.Iterator[T], keyBy fx.Func[T, K]) *distinctArray[T, K] {
	d := &distinctArray[T, K]{
		iter:     iter,
		m:        make(map[K]bool),
		distinct: keyBy,
	}
	d.commonArrayOps = newCommonArrayOps[T](d)
	return d
}

type distinctArray[T any, K comparable] struct {
	iter     types.Iterator[T]
	m        map[K]bool
	distinct fx.Func[T, K]
	commonArrayOps[T]
}

func (d distinctArray[T, K]) Next() (T, bool) {
	for {
		v, ok := d.iter.Next()
		if !ok {
			return v, false
		}
		k := d.distinct(v)
		if !d.m[k] {
			d.m[k] = true
			return v, true
		}
	}
}
