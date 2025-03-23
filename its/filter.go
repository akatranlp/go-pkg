package its

import (
	"iter"
)

type Predicate[T any] = func(T) bool
type Predicate2[K, V any] = func(K, V) bool

func Filter[T any](seq iter.Seq[T], predicate Predicate[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if !predicate(v) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

func Filter2[K, V any](seq iter.Seq2[K, V], predicate Predicate2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !predicate(k, v) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

func All[T any](seq iter.Seq[T], predicate Predicate[T]) bool {
	for v := range seq {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func All2[K, V any](seq iter.Seq2[K, V], predicate Predicate2[K, V]) bool {
	for k, v := range seq {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

func Any[T any](seq iter.Seq[T], predicate Predicate[T]) bool {
	for v := range seq {
		if predicate(v) {
			return true
		}
	}
	return false
}

func Any2[K, V any](seq iter.Seq2[K, V], predicate Predicate2[K, V]) bool {
	for k, v := range seq {
		if predicate(k, v) {
			return true
		}
	}
	return false
}
