package its

import "iter"

func Window[T any](seq iter.Seq[T], n int) iter.Seq[[]T] {
	if n < 2 {
		panic("windowsize is less then 2")
	}

	window := make([]T, 0)
	return func(yield func([]T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for range n {
			v, ok := next()
			if !ok {
				yield(window)
				return
			}
			window = append(window, v)
		}
		ip := pullFromIter(next)
		for ip.Next() {
			window = append(window[1:], ip.Value())
			if !yield(window[1:len(window):len(window)]) {
				return
			}
		}
	}
}

func Window2[T any](seq iter.Seq[T]) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		var first, second T
		var ok bool
		second, ok = next()
		if !ok {
			return
		}

		ip := pullFromIter(next)
		for ip.Next() {
			first = second
			second = ip.Value()
			if !yield(first, second) {
				return
			}
		}
	}
}
