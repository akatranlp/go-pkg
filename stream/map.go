package stream

import (
	"iter"
)

type MapFunc[T, U any] = func(T) U
type Map12Func[T, K, V any] = func(T) (K, V)
type Map21Func[K, V, T any] = func(K, V) T
type Map22Func[T, U, K, V any] = func(T, U) (K, V)

func Map[T, U any](s Stream[T], fn MapFunc[T, U]) Stream[U] {
	return func(yield func(U) bool) {
		for v := range s {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

func Map12[T, U, V any](s Stream[T], fn Map12Func[T, U, V]) Stream2[U, V] {
	return func(yield func(U, V) bool) {
		for v := range s {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

func Map21[T, U, V any](s Stream2[T, U], fn Map21Func[T, U, V]) Stream[V] {
	return func(yield func(V) bool) {
		for k, v := range s {
			if !yield(fn(k, v)) {
				return
			}
		}
	}
}

func Map22[T, U, V, W any](s Stream2[T, U], fn Map22Func[T, U, V, W]) Stream2[V, W] {
	return func(yield func(V, W) bool) {
		for k, v := range s {
			if !yield(fn(k, v)) {
				return
			}
		}
	}
}

func FlatMap[T, U any](s Stream[T], fn MapFunc[T, iter.Seq[U]]) Stream[U] {
	return func(yield func(U) bool) {
		for v := range s {
			for v := range fn(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func FlatMap12[T, U, V any](s Stream[T], fn MapFunc[T, iter.Seq2[U, V]]) Stream2[U, V] {
	return func(yield func(U, V) bool) {
		for v := range s {
			for k, v := range fn(v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

func FlatMap21[T, U, V any](s Stream2[T, U], fn Map21Func[T, U, iter.Seq[V]]) Stream[V] {
	return func(yield func(V) bool) {
		for k, v := range s {
			for v := range fn(k, v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func FlatMap22[T, U, V, W any](s Stream2[T, U], fn Map21Func[T, U, iter.Seq2[V, W]]) Stream2[V, W] {
	return func(yield func(V, W) bool) {
		for k, v := range s {
			for k, v := range fn(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
