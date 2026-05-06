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

func ZipSlices[T, K any](slice1 []T, slice2 []K) iter.Seq2[T, K] {
	length := max(len(slice1), len(slice2))
	return func(yield func(T, K) bool) {
		for i := range length {
			if !yield(slice1[i], slice2[i]) {
				return
			}
		}
	}
}
