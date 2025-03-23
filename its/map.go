package its

import (
	"iter"
)

type MapFunc[T, U any] = func(T) U
type Map12Func[T, K, V any] = func(T) (K, V)
type Map21Func[K, V, T any] = func(K, V) T
type Map22Func[T, U, K, V any] = func(T, U) (K, V)

func Map[T, U any](seq iter.Seq[T], fn MapFunc[T, U]) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range seq {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

func Map12[T, K, V any](seq iter.Seq[T], fn Map12Func[T, K, V]) iter.Seq2[K, V] {
	return func(yield func(k K, v V) bool) {
		for v := range seq {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

func Map21[K, V, T any](seq iter.Seq2[K, V], fn Map21Func[K, V, T]) iter.Seq[T] {
	return func(yield func(v T) bool) {
		for k, v := range seq {
			if !yield(fn(k, v)) {
				return
			}
		}
	}
}

func Map22[T, U, K, V any](seq iter.Seq2[T, U], fn Map22Func[T, U, K, V]) iter.Seq2[K, V] {
	return func(yield func(k K, v V) bool) {
		for k, v := range seq {
			if !yield(fn(k, v)) {
				return
			}
		}
	}
}
