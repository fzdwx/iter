package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

func Distinct[T any, K comparable](
	iter types.Iterator[T],
	keyBy fx.Func[T, K],
) *distinctStream[T, K] {
	d := &distinctStream[T, K]{
		iter:     iter,
		m:        make(map[K]bool),
		distinct: keyBy,
	}
	d.commonStreamOps = newCommonArrayOps[T](d)
	return d
}

type distinctStream[T any, K comparable] struct {
	iter     types.Iterator[T]
	m        map[K]bool
	distinct fx.Func[T, K]
	commonStreamOps[T]
}

func (d distinctStream[T, K]) Next() (T, bool) {
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
