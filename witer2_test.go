package witer_test

import (
	"maps"
	"testing"

	"github.com/hirokisan/witer"
	"github.com/test-go/testify/assert"
)

func Test_witer2(t *testing.T) {
	t.Run("filter", func(t *testing.T) {
		iter := witer.New2(maps.All(map[int64]string{
			0: "a",
			1: "b",
			2: "c",
		}))
		got := maps.Collect(iter.Filter(func(k int64, v string) bool {
			return k%2 == 0 && v == "a"
		}).Iter())
		want := map[int64]string{
			0: "a",
		}
		assert.Equal(t, want, got)
	})

	t.Run("apply", func(t *testing.T) {
		iter := witer.New2(maps.All(map[int64]string{
			0: "a",
			1: "b",
			2: "c",
		}))
		got := maps.Collect(iter.Apply(func(k int64, v string) (int64, string) {
			return k * 10, v
		}).Iter())
		want := map[int64]string{
			0:  "a",
			10: "b",
			20: "c",
		}
		assert.Equal(t, want, got)
	})
}
