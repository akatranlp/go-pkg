package opt

import (
	"iter"
)

func Iter[T any](in Optional[T]) iter.Seq[T] {
	if in.ok {
		return func(yield func(T) bool) {
			yield(in.value)
		}
	}
	return func(yield func(T) bool) {}
}
