package seq

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/stream"
)

func Map[T, R any](seq Seq[T], apply fx.Func[T, R]) Seq[R] {
	return Generator[R](func(accept fx.Consumer[R]) {
		seq.Consume(func(t T) {
			accept(apply(t))
		})
	})
}

func FlatMap[T, R any](seq Seq[T], apply fx.Func[T, Seq[R]]) Seq[R] {
	return Generator[R](func(accept fx.Consumer[R]) {
		seq.Consume(func(t T) {
			apply(t).Consume(accept)
		})
	})
}

func Filter[T any](seq Seq[T], filter fx.Predicate[T]) Seq[T] {
	return Generator[T](func(accept fx.Consumer[T]) {
		seq.Consume(func(t T) {
			if filter(t) {
				accept(t)
			}
		})
	})
}

func Take[T any](seq Seq[T], n int) Seq[T] {
	return Generator[T](func(accept fx.Consumer[T]) {
		if n <= 0 {
			return
		}
		defer func() {
			if err := recover(); err == StopError {
				return
			} else {
				panic(err)
			}
		}()
		seq.Consume(func(v T) {
			accept(v)
			n--
			if n <= 0 {
				panic(StopError) // stop
			}
		})
	})
}

func Drop[T any](seq Seq[T], n int) Seq[T] {
	return Generator[T](func(accept fx.Consumer[T]) {
		seq.Consume(func(v T) {
			if n <= 0 {
				accept(v)
			} else {
				n--
			}
		})
	})
}

func Zip[T, E, R any](seq Seq[T], iter stream.Iterator[E], zipper fx.BiFunc[T, E, R]) Seq[R] {
	return Generator[R](func(accept fx.Consumer[R]) {
		defer func() {
			if err := recover(); err == StopError {
				return
			} else {
				panic(err)
			}
		}()
		if next, ok := iter.Next(); ok {
			seq.Consume(func(t T) {
				accept(zipper(t, next))
			})
		} else {
			panic(StopError)
		}
	})
}

func OnEach[T any](seq Seq[T], consumer fx.Consumer[T]) Seq[T] {
	return Generator[T](func(accept fx.Consumer[T]) {
		seq.Consume(func(t T) {
			consumer(t)
			accept(t)
		})
	})
}
