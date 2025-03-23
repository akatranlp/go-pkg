package its_test

import (
	"errors"
	"slices"
	"strconv"
	"testing"

	"github.com/akatranlp/go-pkg/its"
	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	t.Run("test empty iter", func(t *testing.T) {
		// given
		slice := []int{}
		seq := slices.Values(slice)
		var expected []string

		// when
		actual := its.Map(seq, strconv.Itoa)
		actualSlice := slices.Collect(actual)

		// then
		assert.Equal(t, len(expected), len(actualSlice))
		assert.Equal(t, expected, actualSlice)
	})

	t.Run("test non empty iter", func(t *testing.T) {
		// given
		slice := []int{1, 2, 3}
		seq := slices.Values(slice)
		expected := []string{"1", "2", "3"}

		// when
		actual := its.Map(seq, strconv.Itoa)
		actualSlice := slices.Collect(actual)

		// then
		assert.Equal(t, len(expected), len(actualSlice))
		assert.Equal(t, expected, actualSlice)
	})
}

type testIts struct {
	numberElements int
	called         int
	stopped        bool
}

var _ its.Its[int] = (*testIts)(nil)

func (i *testIts) HasNext() bool {
	return i.called < i.numberElements
}

func (i *testIts) Next() int {
	i.called++
	return i.called
}

func (i *testIts) Close() error {
	if i.stopped {
		return errors.New("already stopped")
	}
	i.stopped = true
	return nil
}

func Test_Its(t *testing.T) {
	t.Run("test empty interface", func(t *testing.T) {
		// given
		testIter := &testIts{numberElements: 0}

		// when
		actual := its.From(testIter)
		empty := its.Empty(actual)

		// then
		assert.True(t, empty)
		assert.Equal(t, 0, testIter.called)
		assert.True(t, testIter.stopped)
	})

	t.Run("test non empty interface", func(t *testing.T) {
		// given
		testIter := testIts{numberElements: 1}
		testIter2 := testIter

		// when
		actual := its.From(&testIter)
		actualSlice := slices.Collect(actual)
		actual = its.From(&testIter2)
		empty := its.Empty(actual)

		// then
		assert.False(t, empty)
		assert.Len(t, actualSlice, 1)
		assert.Equal(t, 1, testIter.called)
		assert.True(t, testIter.stopped)
	})
}
