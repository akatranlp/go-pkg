package its

import "iter"

func Zip[T, U any](seq1 iter.Seq[T], seq2 iter.Seq[U]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()
		puller1 := pullFromIter(next1)

		next2, stop2 := iter.Pull(seq2)
		defer stop2()
		puller2 := pullFromIter(next2)

		for puller1.Next() && puller2.Next() {
			if !yield(puller1.Value(), puller2.Value()) {
				return
			}
		}
	}
}
