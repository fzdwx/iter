package seq

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/stream"
)

func Map[T, R any](seq Seq[T], mapper fx.Func[T, R]) Seq[R] {
	return newimpl[R](&mapSeq[T, R]{seq: seq, mapper: mapper})
}

func FlatMap[T, R any](seq Seq[T], mapper fx.Func[T, Seq[R]]) Seq[R] {
	return newimpl[R](&flatMapSeq[T, R]{seq: seq, mapper: mapper})
}

func Filter[T any](seq Seq[T], filter fx.Predicate[T]) Seq[T] {
	return newimpl[T](&filterSeq[T]{seq: seq, filter: filter})
}

func Take[T any](seq Seq[T], n int) Seq[T] {
	return newimpl[T](&takeSeq[T]{seq: seq, n: n})
}

func Drop[T any](seq Seq[T], n int) Seq[T] {
	return newimpl[T](&dropSeq[T]{seq: seq, n: n})
}

func Zip[T, E, R any](seq Seq[T], iter stream.Iterator[E], zipper fx.BiFunc[T, E, R]) Seq[R] {
	return newimpl[R](&zipSeq[T, E, R]{seq: seq, iter: iter, mapper: zipper})
}

func OnEach[T any](seq Seq[T], consumer fx.Consumer[T]) Seq[T] {
	return &onEachSeq[T]{seq: seq, fx: consumer}
}
