package set

import (
	"github.com/akatranlp/go-pkg/stream"
	"iter"
)

type Set[T comparable] map[T]struct{}

func New[T comparable](values ...T) Set[T] {
	s := make(Set[T])
	s.SetValues(values...)
	return s
}

func Collect[T comparable](seq iter.Seq[T]) Set[T] {
	s := make(Set[T])
	s.SetIter(seq)
	return s
}

func (s Set[T]) Length() int {
	return len(s)
}

func (s Set[T]) Empty() bool {
	return len(s) == 0
}

func (s Set[T]) Set(k T) bool {
	_, ok := s[k]
	s[k] = struct{}{}
	return !ok
}

func (s Set[T]) SetIter(seq iter.Seq[T]) {
	for k := range seq {
		s[k] = struct{}{}
	}
}

func (s Set[T]) SetValues(slice ...T) {
	for _, k := range slice {
		s[k] = struct{}{}
	}
}

func (s Set[T]) Has(k T) bool {
	_, ok := s[k]
	return ok
}

func (s Set[T]) HasAll(seq iter.Seq[T]) bool {
	for k := range seq {
		_, ok := s[k]
		if !ok {
			return false
		}
	}
	return true
}

func (s Set[T]) Remove(k T) bool {
	_, ok := s[k]
	delete(s, k)
	return ok
}

func (s Set[T]) Outer(other Set[T]) Set[T] {
	result := make(Set[T])
	for k, v := range s {
		if _, ok := other[k]; !ok {
			result[k] = v
		}
	}
	for k, v := range other {
		if _, ok := s[k]; !ok {
			result[k] = v
		}
	}
	return result
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	result := make(Set[T])
	for k, v := range s {
		if _, ok := other[k]; ok {
			result[k] = v
		}
	}
	return result
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := make(Set[T])
	for k, v := range s {
		result[k] = v
	}
	for k, v := range other {
		result[k] = v
	}
	return result
}

func (s Set[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range s {
			if !yield(k) {
				return
			}
		}
	}
}

func (s Set[T]) Stream() stream.Stream[T] {
	return stream.Of(s.Iter())
}
