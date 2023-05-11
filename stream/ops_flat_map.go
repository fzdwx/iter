package stream

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/types"
)

func FlatMap[T, R any](iter Iterator[T], mapper fx.Func[T, Iterator[R]]) *flatMapStream[T, R] {
	f := &flatMapStream[T, R]{
		mapper: mapper,
		iter:   iter,
	}
	f.commonStreamOps = newCommonArrayOps[R](f)
	return f
}

type flatMapStream[T, R any] struct {
	iter   Iterator[T]
	mapper fx.Func[T, Iterator[R]]
	commonStreamOps[R]

	currentIter Iterator[R]
}

func (f *flatMapStream[T, R]) Next() (R, bool) {
	if f.currentIter == nil {
		if nextT, ok := f.iter.Next(); !ok {
			return types.Empty[R](), ok
		} else {
			f.currentIter = f.mapper(nextT)
		}
	}

	if nextR, ok := f.currentIter.Next(); ok {
		return nextR, ok
	}

	f.currentIter = nil
	return f.Next()
}
