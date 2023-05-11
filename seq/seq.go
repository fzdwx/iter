package seq

import (
	"errors"
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/stream"
)

var StopError = errors.New("seq stop")

type Seq[T any] interface {
	Consume(accept fx.Consumer[T])
}

type unit[T any] struct {
	t T
}

func (u *unit[T]) Consume(accept fx.Consumer[T]) {
	accept(u.t)
}

type iterSeq[T any] struct {
	iter stream.Iterator[T]
}

func (i *iterSeq[T]) Consume(accept fx.Consumer[T]) {
	stream.ForEach(i.iter, accept)
}

type mapSeq[T, R any] struct {
	seq    Seq[T]
	mapper fx.Func[T, R]
}

func (m *mapSeq[T, R]) Consume(accept fx.Consumer[R]) {
	m.seq.Consume(func(t T) {
		accept(m.mapper(t))
	})
}

type flatMapSeq[T, R any] struct {
	seq    Seq[T]
	mapper fx.Func[T, Seq[R]]
}

func (m *flatMapSeq[T, R]) Consume(accept fx.Consumer[R]) {
	m.seq.Consume(func(t T) {
		m.mapper(t).Consume(accept)
	})
}

type filterSeq[T any] struct {
	seq    Seq[T]
	filter fx.Predicate[T]
}

func (f *filterSeq[T]) Consume(accept fx.Consumer[T]) {
	f.seq.Consume(func(t T) {
		if f.filter(t) {
			accept(t)
		}
	})
}

type takeSeq[T any] struct {
	seq Seq[T]
	n   int
}

func (t *takeSeq[T]) Consume(accept fx.Consumer[T]) {
	if t.n <= 0 {
		return
	}
	defer func() {
		if err := recover(); err == StopError {
			return
		} else {
			panic(err)
		}
	}()
	t.seq.Consume(func(v T) {
		accept(v)
		t.n--
		if t.n <= 0 {
			panic(StopError) // stop
		}
	})
}

type dropSeq[T any] struct {
	seq Seq[T]
	n   int
}

func (t *dropSeq[T]) Consume(accept fx.Consumer[T]) {
	t.seq.Consume(func(v T) {
		if t.n <= 0 {
			accept(v)
		} else {
			t.n--
		}
	})
}

type zipSeq[T, E, R any] struct {
	iter   stream.Iterator[E]
	mapper fx.BiFunc[T, E, R]
	seq    Seq[T]
}

func (z *zipSeq[T, E, R]) Consume(accept fx.Consumer[R]) {
	defer func() {
		if err := recover(); err == StopError {
			return
		} else {
			panic(err)
		}
	}()
	if next, ok := z.iter.Next(); ok {
		z.seq.Consume(func(t T) {
			accept(z.mapper(t, next))
		})
	} else {
		panic(StopError)
	}
}

type onEachSeq[T any] struct {
	seq Seq[T]
	fx  fx.Consumer[T]
}

func (o *onEachSeq[T]) Consume(accept fx.Consumer[T]) {
	o.seq.Consume(func(t T) {
		accept(t)
		o.fx(t)
	})
}

type quickSeq[T any] struct {
	f func(accept fx.Consumer[T])
}

func (q *quickSeq[T]) Consume(accept fx.Consumer[T]) {
	q.f(accept)
}

func (q *quickSeq[T]) build(f func(accept fx.Consumer[T])) {
	q.f = f
}
