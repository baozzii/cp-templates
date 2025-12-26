package common

import (
	"cmp"
	"fmt"
	"strconv"
)

func ToString[T any](e T) string {
	return fmt.Sprintf("%v", e)
}

func ToInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func Chmax[T cmp.Ordered](x *T, y T) {
	*x = max(*x, y)
}

func Chmin[T cmp.Ordered](x *T, y T) {
	*x = min(*x, y)
}

func Sum[T ComplexNumber, E ~[]T](v E) T {
	var s T
	for _, w := range v {
		s += w
	}
	return s
}

func PreSum[T ComplexNumber, E ~[]T](v E) E {
	p := make(E, len(v)+1)
	for i, w := range v {
		p[i+1] = p[i] + w
	}
	return p
}

func Count[T comparable, E ~[]T](v E, e T) int {
	cnt := 0
	for _, w := range v {
		if w == e {
			cnt++
		}
	}
	return cnt
}

func Iota[T Integer, E ~[]T](v E, e T) {
	for i := range v {
		v[i] = e + T(i)
	}
}
