package witer

import (
	"iter"
)

type witer2[K, V any] struct {
	inner iter.Seq2[K, V]
}

func New2[K, V any](inner iter.Seq2[K, V]) *witer2[K, V] {
	return &witer2[K, V]{inner}
}

func (i *witer2[K, V]) Filter(f func(K, V) bool) *witer2[K, V] {
	inner := func(yield func(K, V) bool) {
		for k, v := range i.inner {
			if f(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
	return &witer2[K, V]{inner}
}

func (i *witer2[K, V]) Apply(f func(K, V) (K, V)) *witer2[K, V] {
	inner := func(yield func(K, V) bool) {
		for k, v := range i.inner {
			if !yield(f(k, v)) {
				return
			}
		}
	}
	return &witer2[K, V]{inner}
}

func (i *witer2[K, V]) Iter() iter.Seq2[K, V] {
	return i.inner
}
