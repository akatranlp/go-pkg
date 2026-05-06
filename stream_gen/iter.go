package main

import (
	"io"
	"iter"
)

type Its[T any] interface {
	Next() (T, bool)
}

func From[T any](i Its[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		if c, ok := i.(io.Closer); ok {
			defer c.Close()
		}
		for {
			v, ok := i.Next()
			if !ok {
				return
			} else if !yield(v) {
				return
			}
		}
	}
}

func Zip[T, U any](seq1 iter.Seq[T], seq2 iter.Seq[U]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()

		next2, stop2 := iter.Pull(seq2)
		defer stop2()

		for {
			v1, ok1 := next1()
			v2, ok2 := next2()
			if !ok1 || !ok2 {
				return
			} else if !yield(v1, v2) {
				return
			}
		}
	}
}

func Zip_Longest[T, U any](seq1 iter.Seq[T], seq2 iter.Seq[U]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()

		next2, stop2 := iter.Pull(seq2)
		defer stop2()

		for {
			v1, ok1 := next1()
			v2, ok2 := next2()
			if !ok1 && !ok2 {
				return
			} else if !yield(v1, v2) {
				return
			}
		}
	}
}

func Zip_Strict[T, U any](seq1 iter.Seq[T], seq2 iter.Seq[U]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()

		next2, stop2 := iter.Pull(seq2)
		defer stop2()

		for {
			v1, ok1 := next1()
			v2, ok2 := next2()
			if !ok1 && ok2 || ok1 && !ok2 {
				panic("not equal length")
			} else if !ok1 && !ok2 {
				return
			} else if !yield(v1, v2) {
				return
			}
		}
	}
}
