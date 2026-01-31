package templates

import (
	"slices"
)

type vector[T any] []T

func (v *vector[T]) size() int {
	return len(*v)
}

func (v *vector[T]) empty() bool {
	return len(*v) == 0
}

func (v *vector[T]) clear() {
	*v = (*v)[:0]
}

func (v *vector[T]) push_back(x T) {
	*v = append(*v, x)
}

func (v *vector[T]) pop_back() {
	*v = (*v)[:len(*v)-1]
}

func (v *vector[T]) resize(x int) {
	if v.size() > x {
		*v = (*v)[:x]
	} else {
		*v = append(*v, make([]T, x-v.size())...)
	}
}

func (v *vector[T]) back() *T {
	return &((*v)[v.size()-1])
}

func (v *vector[T]) front() *T {
	return &((*v)[0])
}

func (v *vector[T]) erase(x, y int) {
	*v = slices.Delete(*v, x, y)
}

func (v *vector[T]) insert(x int, w ...T) {
	*v = slices.Insert(*v, x, w...)
}

func (v *vector[T]) reverse() {
	n := v.size()
	for i := 0; i < n/2; i++ {
		(*v)[i], (*v)[n-i-1] = (*v)[n-i-1], (*v)[i]
	}
}

func (v *vector[T]) fill(w T) {
	for i := range *v {
		(*v)[i] = w
	}
}

func (v *vector[T]) copy() vector[T] {
	return slices.Clone(*v)
}

func (v *vector[T]) slice(i, j int) vector[T] {
	w := (*v)[i:j]
	return w.copy()
}

func (v *vector[T]) scan(read func(...any)) {
	for _, w := range *v {
		read(w)
	}
}

func (v *vector[T]) concat(w vector[T]) {
	for _, c := range w {
		v.push_back(c)
	}
}

func concat[T any](a, b vector[T]) vector[T] {
	na := a.copy()
	na.concat(b)
	return na
}

func vec1[T any, S integer](n ...S) vector[T] {
	if len(n) == 0 {
		return make(vector[T], 0)
	} else {
		return make(vector[T], n[0])
	}
}

func vec2[T any, S integer](n ...S) vector[vector[T]] {
	if len(n) == 0 {
		return make(vector[vector[T]], 0)
	} else {
		res := make(vector[vector[T]], n[0])
		for i := range res {
			res[i] = vec1[T](n[1:]...)
		}
		return res
	}
}

func vec3[T any, S integer](n ...S) vector[vector[vector[T]]] {
	if len(n) == 0 {
		return make(vector[vector[vector[T]]], 0)
	} else {
		res := make(vector[vector[vector[T]]], n[0])
		for i := range res {
			res[i] = vec2[T](n[1:]...)
		}
		return res
	}
}

func to_umap[T comparable, E ~[]T](v E) umap[T, int] {
	cnt := make(umap[T, int])
	for _, w := range v {
		cnt[w]++
	}
	return cnt
}

func to_uset[T comparable, E ~[]T](v E) uset[T] {
	cnt := make(uset[T])
	for _, w := range v {
		cnt.insert(w)
	}
	return cnt
}
