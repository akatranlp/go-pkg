package its

import (
	"errors"
	"iter"
)

func CollectOrError[T any](seq iter.Seq2[T, error]) ([]T, error) {
	var result []T
	for v, err := range seq {
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func CollectOrJoinError[T any](seq iter.Seq2[T, error]) (result []T, err error) {
	for v, newErr := range seq {
		if err != nil {
			err = errors.Join(err, newErr)
		} else {
			result = append(result, v)
		}
	}
	return
}
