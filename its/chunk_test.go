package its_test

import (
	"github.com/akatranlp/go-pkg/its"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test_Chunk(t *testing.T) {
	t.Run("test empty iter", func(t *testing.T) {
		// given
		slice := []int{}
		seq := slices.Values(slice)
		chunksize := 2

		// when
		actual := its.Chunk(seq, chunksize)
		empty := its.Empty(actual)

		// then
		assert.True(t, empty)
	})

	t.Run("test non empty iter less then chunksize", func(t *testing.T) {
		slice := []int{1}
		seq := slices.Values(slice)
		chunksize := 2

		// when
		actual := its.Chunk(seq, chunksize)
		empty := its.Empty(actual)

		var called int
		its.Foreach(actual, func(v []int) {
			called++
			assert.Len(t, v, 1)
		})

		// then
		assert.False(t, empty)
		assert.Equal(t, 1, called)
	})

	t.Run("test non empty iter with same size as chunk", func(t *testing.T) {
		slice := []int{1, 2}
		seq := slices.Values(slice)
		chunksize := 2

		// when
		actual := its.Chunk(seq, chunksize)
		empty := its.Empty(actual)

		var called int
		its.Foreach(actual, func(v []int) {
			called++
			assert.Len(t, v, 2)
		})

		// then
		assert.False(t, empty)
		assert.Equal(t, 1, called)
	})

	t.Run("test non empty iter more than chunksize", func(t *testing.T) {
		slice := []int{1, 2, 3}
		seq := slices.Values(slice)
		chunksize := 2

		// when
		actual := its.Chunk(seq, chunksize)
		empty := its.Empty(actual)

		var called int
		its.Foreach2(its.Enumerate(actual), func(idx int, v []int) {
			called++
			if idx == 0 {
				assert.Len(t, v, 2)
			}
			if idx == 1 {
				assert.Len(t, v, 1)
			}
		})

		// then
		assert.False(t, empty)
		assert.Equal(t, 2, called)
	})

	t.Run("test chunksize panic", func(t *testing.T) {
		its.Foreach(its.Range(3), func(i int) {
			// given
			slice := []int{}
			seq := slices.Values(slice)
			chunksize := i - 1

			defer func() {
				assert.NotNil(t, recover())
			}()

			// when
			its.Chunk(seq, chunksize)
		})
	})
}

func Test_Chunk2(t *testing.T) {
	t.Run("test empty iter", func(t *testing.T) {
		// given
		slice := []int{}
		seq := slices.Values(slice)

		// when
		actual := its.Chunk2(seq)
		empty := its.Empty2(actual)

		// then
		assert.True(t, empty)
	})

	t.Run("test non empty iter", func(t *testing.T) {
		slice := []int{1}
		seq := slices.Values(slice)

		var called int
		// when
		actual := its.Chunk2(seq)
		empty := its.Empty2(actual)
		its.Foreach2(actual, func(v1, v2 int) {
			called++
			assert.Equal(t, 1, v1)
			assert.Equal(t, 0, v2)
		})

		// then
		assert.False(t, empty)
		assert.Equal(t, 1, called)
	})
}
