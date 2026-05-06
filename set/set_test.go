package set

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Set(t *testing.T) {
	t.Run("test empty set", func(t *testing.T) {
		// given
		s := New[int]()

		// when
		empty := s.Empty()
		slice := slices.Collect(s.Iter())

		// then
		assert.True(t, empty)
		assert.Equal(t, 0, len(slice))
	})

	t.Run("merge set", func(t *testing.T) {
		// given
		s1 := NewWithValues(1, 2, 3)
		s2 := NewWithValues(3, 4, 5)

		// when
		s3 := NewFromIter(s1.Iter()).SetIter(s2.Iter())

		allS1 := s1.Stream().All(func(v int) bool { return s3.Has(v) })
		allS2 := s2.Stream().All(func(v int) bool { return s3.Has(v) })
		length := s3.Length()

		// then
		assert.True(t, allS1)
		assert.True(t, allS2)
		assert.Equal(t, 5, length)
	})
}
