package seq

import (
	"github.com/fzdwx/iter/fx"
	"strings"
)

func newimpl[T any](seq Seq[T]) *Impl[T] {
	return &Impl[T]{seq: seq}
}

type Impl[T any] struct {
	seq Seq[T]
}

func (d *Impl[T]) Filter(filter fx.Predicate[T]) *Impl[T] {
	s := Filter(d.seq, filter)
	d.seq = s
	return d
}

func (d *Impl[T]) Take(n int) *Impl[T] {
	s := Take(d.seq, n)
	d.seq = s
	return d
}

func (d *Impl[T]) Drop(n int) *Impl[T] {
	s := Drop(d.seq, n)
	d.seq = s
	return d
}

func (d *Impl[T]) OnEach(accept fx.Consumer[T]) {
	d.seq.Consume(accept)
}

func (d *Impl[T]) Consume(accept fx.Consumer[T]) {
	d.seq.Consume(accept)
}

func (d *Impl[T]) Join(sep string, mapper fx.Func[T, string]) string {
	var sb strings.Builder
	d.seq.Consume(func(t T) {
		sb.WriteString(mapper(t))
		sb.WriteString(sep)
	})

	res := sb.String()
	return strings.TrimSuffix(res, sep)
}

func (d *Impl[T]) Collect() []T {
	result := make([]T, 0)
	d.seq.Consume(func(t T) {
		result = append(result, t)
	})
	return result
}
