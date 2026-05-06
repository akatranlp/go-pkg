package stream

import (
	"iter"
	"maps"
	"slices"

	"github.com/akatranlp/go-pkg/opt"
)

type Stream[T any] iter.Seq[T]
type Stream2[T, U any] iter.Seq2[T, U]

func Of[T any](seq iter.Seq[T]) Stream[T] {
	return Stream[T](seq)
}

func Of2[T, U any](seq iter.Seq2[T, U]) Stream2[T, U] {
	return Stream2[T, U](seq)
}

func (s Stream[T]) To() iter.Seq[T] {
	return iter.Seq[T](s)
}

func (s Stream2[T, U]) To() iter.Seq2[T, U] {
	return iter.Seq2[T, U](s)
}

func (s Stream[T]) FindFirst() opt.Optional[T] {
	for v := range s {
		return opt.Some(v)
	}
	return opt.None[T]()
}

func (s Stream[T]) CollectSlice() []T {
	return slices.Collect(s.To())
}

func CollectMap[T any, U comparable, V any](s Stream[T], fn Map12Func[T, U, V]) map[U]V {
	return maps.Collect(Map12(s, fn).To())
}

func CollectSet[T any, U comparable](s Stream[T], fn MapFunc[T, U]) map[U]struct{} {
	r := make(map[U]struct{})
	for v := range s {
		r[fn(v)] = struct{}{}
	}
	return r
}

func Id[T any](v T) T {
	return v
}
