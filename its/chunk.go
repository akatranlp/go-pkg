package its

import (
	"iter"
)

func Chunk[T any](seq iter.Seq[T], n int) iter.Seq[[]T] {
	if n < 2 {
		panic("chunksize is less then 2")
	}

	chunk := make([]T, n)
	return func(yield func([]T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()

		for {
			for i := range n {
				v, ok := next()
				if !ok && i == 0 {
					return
				} else if !ok {
					yield(chunk[:i:i])
					return
				}
				chunk[i] = v
			}

			if !yield(chunk[:len(chunk):len(chunk)]) {
				return
			}
		}
	}
}

func Chunk2[T any](seq iter.Seq[T]) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v1, ok1 := next()
			v2, ok2 := next()
			if !ok1 && !ok2 {
				return
			}
			if !yield(v1, v2) {
				return
			}
		}
	}
}
