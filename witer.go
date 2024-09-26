package witer

import (
	"iter"
	"slices"
)

type witer[V any] struct {
	inner iter.Seq[V]
}

func New[V any](inner iter.Seq[V]) *witer[V] {
	return &witer[V]{inner}
}

func (i *witer[V]) Filter(f func(V) bool) *witer[V] {
	inner := func(yield func(V) bool) {
		for v := range i.inner {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
	return &witer[V]{inner}
}

func (i *witer[V]) Apply(f func(V) V) *witer[V] {
	inner := func(yield func(V) bool) {
		for v := range i.inner {
			if !yield(f(v)) {
				return
			}
		}
	}
	return &witer[V]{inner}
}

func (i *witer[V]) Iter() iter.Seq[V] {
	return i.inner
}

func (i *witer[V]) Collect() []V {
	return slices.Collect(i.inner)
}
