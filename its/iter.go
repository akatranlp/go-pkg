package its

import (
	"io"
	"iter"
)

func Foreach[T any](seq iter.Seq[T], fn func(T)) {
	for v := range seq {
		fn(v)
	}
}

func Foreach2[K, V any](seq iter.Seq2[K, V], fn func(K, V)) {
	for k, v := range seq {
		fn(k, v)
	}
}

func Enumerate[T any](seq iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		var idx int
		for v := range seq {
			if !yield(idx, v) {
				return
			}
			idx++
		}
	}
}

func Range(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range n {
			if !yield(i) {
				return
			}
		}
	}
}

func Empty[T any](seq iter.Seq[T]) bool {
	next, stop := iter.Pull(seq)
	defer stop()
	_, ok := next()
	return !ok
}

func Empty2[K, V any](seq iter.Seq2[K, V]) bool {
	next, stop := iter.Pull2(seq)
	defer stop()
	_, _, ok := next()
	return !ok
}

type iterPuller[T any] struct {
	next  func() (T, bool)
	value *T
}

func (i *iterPuller[T]) Next() bool {
	v, ok := i.next()
	i.value = &v
	return ok
}

func (i *iterPuller[T]) Value() T {
	return *i.value
}

func pullFromIter[T any](next func() (T, bool)) *iterPuller[T] {
	return &iterPuller[T]{next: next}
}

type Its[T any] interface {
	HasNext() bool
	Next() T
}

type Its2[K, V any] interface {
	HasNext() bool
	Next() (K, V)
}

/*** From
* If you implement io.Closer then the iterator will be closed in the end
 */
func From[T any](its Its[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		if closer, ok := its.(io.Closer); ok {
			defer closer.Close()
		}
		for its.HasNext() {
			if !yield(its.Next()) {
				return
			}
		}
	}
}

func From2[K, V any](its Its2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if closer, ok := its.(io.Closer); ok {
			defer closer.Close()
		}
		for its.HasNext() {
			if !yield(its.Next()) {
				return
			}
		}
	}
}
