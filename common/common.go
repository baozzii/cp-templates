package common

import "cmp"

func Chmax[T cmp.Ordered](x *T, y T) {
	*x = max(*x, y)
}

func Chmin[T cmp.Ordered](x *T, y T) {
	*x = min(*x, y)
}

func NewSlice1[T any](n int) []T {
	return make([]T, n)
}

func NewSlice2[T any](n, m int) [][]T {
	r := make([][]T, n)
	for i := range r {
		r[i] = make([]T, m)
	}
	return r
}

func NewSlice3[T any](n, m, k int) [][][]T {
	r := make([][][]T, n)
	for i := range r {
		r[i] = make([][]T, m)
		for j := range r[i] {
			r[i][j] = make([]T, k)
		}
	}
	return r
}

func Fill[T any](a []T, v T) {
	for i := range a {
		a[i] = v
	}
}
