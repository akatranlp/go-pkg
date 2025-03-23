package its_test

import (
	"errors"
	"iter"
	"slices"
	"testing"

	"github.com/akatranlp/go-pkg/its"
	"github.com/stretchr/testify/assert"
)

var testError = errors.New("test error")

func errorSeq(slice []int, errIdx int) iter.Seq2[int, error] {
	return its.Map22(slices.All(slice), func(i, v int) (int, error) {
		if i == errIdx {
			return 0, testError
		}
		return v, nil
	})
}

func Test_Error_Map(t *testing.T) {
	t.Run("test empty iter", func(t *testing.T) {
		// given
		slice := []int{}
		seq := errorSeq(slice, -1)
		var expected []int

		// when
		actual, err := its.CollectOrError(seq)

		// then
		assert.NoError(t, err)
		assert.Empty(t, actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("test non empty iter no error", func(t *testing.T) {
		// given
		slice := []int{1, 2, 3}
		seq := errorSeq(slice, -1)

		// when
		actual, err := its.CollectOrError(seq)

		// then
		assert.NoError(t, err)
		assert.NotEmpty(t, actual)
		assert.Equal(t, slice, actual)
	})

	t.Run("test non empty iter error", func(t *testing.T) {
		// given
		slice := []int{1, 2, 3}
		seq := errorSeq(slice, 0)
		var expected []int

		// when
		actual, err := its.CollectOrError(seq)

		// then
		assert.Error(t, err)
		assert.ErrorIs(t, testError, err)
		assert.Empty(t, actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("test non empty iter error last index", func(t *testing.T) {
		// given
		slice := []int{1, 2, 3}
		seq := errorSeq(slice, 2)
		var expected []int

		// when
		actual, err := its.CollectOrError(seq)

		// then
		assert.Error(t, err)
		assert.ErrorIs(t, testError, err)
		assert.Empty(t, actual)
		assert.Equal(t, expected, actual)
	})
}
