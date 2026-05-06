package stream

type Predicate[T any] = func(T) bool
type Predicate2[T, U any] = func(T, U) bool

func (s Stream[T]) Filter(predicate Predicate[T]) Stream[T] {
	return func(yield func(T) bool) {
		for v := range s {
			if !predicate(v) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

func (s Stream2[T, U]) Filter(predicate Predicate2[T, U]) Stream2[T, U] {
	return func(yield func(T, U) bool) {
		for k, v := range s {
			if !predicate(k, v) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

func (s Stream[T]) All(predicate Predicate[T]) bool {
	for v := range s {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func (s Stream2[T, U]) All(predicate Predicate2[T, U]) bool {
	for k, v := range s {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

func (s Stream[T]) Any(predicate Predicate[T]) bool {
	for v := range s {
		if predicate(v) {
			return true
		}
	}
	return false
}

func (s Stream2[T, U]) Any(predicate Predicate2[T, U]) bool {
	for k, v := range s {
		if predicate(k, v) {
			return true
		}
	}
	return false
}
