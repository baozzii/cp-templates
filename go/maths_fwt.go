package templates

import "slices"

const (
	FWT_AND = iota
	FWT_OR
	FWT_XOR
)

func fwt[T integer, E ~[]T](a E, f int, t bool) E {
	switch f {
	case FWT_AND:
		a = fwt_and(a, cond(t, 1, -1))
	case FWT_OR:
		a = fwt_or(a, cond(t, 1, -1))
	case FWT_XOR:
		a = fwt_xor(a, cond(t, 1, -1))
	}
	return a
}

func fwt_or[T integer, E ~[]T](a E, t int) E {
	var z T
	n := len(a)
	if _, ok := any(z).(mint); ok {
		b := make([]mint, n)
		for i := range n {
			b[i] = norm(a[i])
		}
		t := norm(t)
		for x := 2; x <= n; x <<= 1 {
			k := x >> 1
			for i := 0; i < n; i += x {
				for j := 0; j < k; j++ {
					b[i+j+k] = b[i+j+k].add(b[i+j].mul(t))
				}
			}
		}
		c := make(E, n)
		for i := range n {
			c[i] = T(b[i])
		}
		return c
	} else {
		b := slices.Clone(a)
		for x := 2; x <= n; x <<= 1 {
			k := x >> 1
			for i := 0; i < n; i += x {
				for j := 0; j < k; j++ {
					b[i+j+k] += b[i+j] * T(t)
				}
			}
		}
		return b
	}
}

func fwt_and[T integer, E ~[]T](a E, t int) E {
	var z T
	n := len(a)
	if _, ok := any(z).(mint); ok {
		b := make([]mint, n)
		for i := range n {
			b[i] = norm(a[i])
		}
		t := norm(t)
		for x := 2; x <= n; x <<= 1 {
			k := x >> 1
			for i := 0; i < n; i += x {
				for j := 0; j < k; j++ {
					b[i+j] = b[i+j].add(b[i+j+k].mul(t))
				}
			}
		}
		c := make(E, n)
		for i := range n {
			c[i] = T(b[i])
		}
		return c
	} else {
		b := slices.Clone(a)
		for x := 2; x <= n; x <<= 1 {
			k := x >> 1
			for i := 0; i < n; i += x {
				for j := 0; j < k; j++ {
					b[i+j] += b[i+j+k] * T(t)
				}
			}
		}
		return b
	}
}

func fwt_xor[T integer, E ~[]T](a E, t int) E {
	var z T
	n := len(a)
	if _, ok := any(z).(mint); ok {
		b := make([]mint, n)
		for i := range n {
			b[i] = norm(a[i])
		}
		for x := 2; x <= n; x <<= 1 {
			k := x >> 1
			for i := 0; i < n; i += x {
				for j := 0; j < k; j++ {
					b[i+j] = b[i+j].add(b[i+j+k])
					b[i+j+k] = b[i+j].sub(b[i+j+k].mul(2))
					if t == -1 {
						b[i+j] = b[i+j].mul((M + 1) / 2)
						b[i+j+k] = b[i+j+k].mul((M + 1) / 2)
					}
				}
			}
		}
		c := make(E, n)
		for i := range n {
			c[i] = T(b[i])
		}
		return c
	} else {
		b := slices.Clone(a)
		for x := 2; x <= n; x <<= 1 {
			k := x >> 1
			for i := 0; i < n; i += x {
				for j := 0; j < k; j++ {
					b[i+j] += b[i+j+k]
					b[i+j+k] = b[i+j] - b[i+j+k]*2
					if t == -1 {
						b[i+j] /= 2
						b[i+j+k] /= 2
					}
				}
			}
		}
		return b
	}
}
