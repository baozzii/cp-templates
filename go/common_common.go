package templates

import (
	"cmp"
	"fmt"
	"strconv"
	"unsafe"
)

type void struct{}

type numeric_limits[T integer] struct{}

func limit[T integer]() numeric_limits[T] {
	return struct{}{}
}

func (numeric_limits[T]) max() T {
	var z T
	if (^z) < 0 {
		b := uint(unsafe.Sizeof(z) * 8)
		u := uint(1)<<(b-1) - 1
		return T(u)
	} else {
		return ^z
	}
}

func (numeric_limits[T]) min() T {
	var z T
	if (^z) < 0 {
		b := uint(unsafe.Sizeof(z) * 8)
		u := uint(1) << (b - 1)
		return T(-int(u))
	}
	return 0
}

func to_string[T any](e T) string {
	return fmt.Sprintf("%v", e)
}

func to_int[T integer](s string) T {
	x, _ := strconv.Atoi(s)
	return T(x)
}

func ckmax[T cmp.Ordered](x *T, y T) {
	*x = max(*x, y)
}

func ckmin[T cmp.Ordered](x *T, y T) {
	*x = min(*x, y)
}

func sum[T complex, E ~[]T](v E) T {
	var s T
	for _, w := range v {
		s += w
	}
	return s
}

func presum[T complex, E ~[]T](v E) E {
	p := make(E, len(v)+1)
	for i, w := range v {
		p[i+1] = p[i] + w
	}
	return p
}

func count[T comparable, E ~[]T](v E, e T) int {
	cnt := 0
	for _, w := range v {
		if w == e {
			cnt++
		}
	}
	return cnt
}

func cond[T any](cond bool, x, y T) T {
	if cond {
		return x
	}
	return y
}
