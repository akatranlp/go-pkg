package its

import "iter"

type ReduceFunc[T, U any] = func(acc U, v T) U
type ReduceWithErrorFunc[T, U any] = func(acc U, v T) (U, error)
type Reduce2Func[K, V, T any] = func(acc T, k K, v V) T
type Reduce2WithErrorFunc[K, V, T any] = func(acc T, k K, v V) (T, error)

func Reduce[T, U any](seq iter.Seq[T], acc U, fn ReduceFunc[T, U]) U {
	for v := range seq {
		acc = fn(acc, v)
	}
	return acc
}

func ReduceWithError[T, U any](seq iter.Seq2[T, error], acc U, fn ReduceWithErrorFunc[T, U]) (U, error) {
	var err error
	for v := range seq {
		acc, err = fn(acc, v)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func Reduce2[K, V, T any](seq iter.Seq2[K, V], acc T, fn Reduce2Func[K, V, T]) T {
	for k, v := range seq {
		acc = fn(acc, k, v)
	}
	return acc
}

func Reduce2WithError[K, V, T any](seq iter.Seq2[K, V], acc T, fn Reduce2WithErrorFunc[K, V, T]) (T, error) {
	var err error
	for k, v := range seq {
		acc, err = fn(acc, k, v)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}
