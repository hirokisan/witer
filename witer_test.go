package witer_test

import (
	"slices"
	"testing"

	"github.com/hirokisan/witer"
	"github.com/test-go/testify/assert"
)

func Test_witer(t *testing.T) {
	t.Run("filter", func(t *testing.T) {
		iter := witer.New(slices.Values([]int64{1, 2, 3}))
		got := iter.Filter(func(v int64) bool {
			return v%2 == 0
		}).Collect()
		want := []int64{2}
		assert.Equal(t, want, got)
	})

	t.Run("apply", func(t *testing.T) {
		iter := witer.New(slices.Values([]int64{1, 2, 3}))
		got := iter.Apply(func(v int64) int64 {
			return v * 10
		}).Collect()
		want := []int64{10, 20, 30}
		assert.Equal(t, want, got)
	})
}
