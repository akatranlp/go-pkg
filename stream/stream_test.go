package stream

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stream(t *testing.T) {
	t.Run("test empty stream", func(t *testing.T) {
		// given
		val := []int{}
		s := Of(slices.Values(val))

		// when
		slice := s.CollectSlice()

		// then
		assert.Equal(t, 0, len(slice))
	})
}

func Test_Collect(t *testing.T) {
	t.Run("slice", func(t *testing.T) {
		// given
		val := []int{1, 2, 3}
		s := Of(slices.Values(val))

		// when
		res := s.CollectSlice()

		// then
		assert.Equal(t, val, res)
	})

	t.Run("map", func(t *testing.T) {
		// given
		val := map[int]int{1: 1, 2: 2, 3: 3}
		s := Of(maps.Values(val))

		// when
		res := CollectMap(s, func(v int) (int, int) { return v, v })

		// then
		assert.Equal(t, val, res)
	})

	t.Run("set", func(t *testing.T) {
		// given
		val := []int{1, 2, 3}
		s := Of(slices.Values(val))

		// when
		res := CollectSet(s, Id)
		resSlice := slices.Collect(maps.Keys(res))

		// then
		assert.Equal(t, val, resSlice)
	})
}
