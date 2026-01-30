package templates

import "math/bits"

const M = 998244353

const (
	__m  = M
	__im = ^uint64(0)/uint64(M) + 1
)

type mint int

func norm[T integer](x T) mint {
	y := int(x)
	y %= M
	if y < 0 {
		y += M
	}
	return mint(y)
}

func (x mint) add(y mint) mint {
	z := x + y
	if z >= M {
		z -= M
	}
	return z
}

func (x mint) sub(y mint) mint {
	z := x - y
	if z < 0 {
		z += M
	}
	return z
}

func (x mint) mul(y mint) mint {
	z := uint64(x * y)
	hi, _ := bits.Mul64(uint64(z), __im)
	p := hi * __m
	r := z - p
	if z < p {
		r += M
	}
	return mint(r)
}

func (x mint) inv() mint {
	_, v, _ := exgcd(x, M)
	if v < 0 {
		v += M
	}
	return v
}

func (x mint) div(y mint) mint {
	return x.mul(y.inv())
}

func (x mint) pow(y int) mint {
	r := norm(1)
	for ; y > 0; x, y = x.mul(x), y>>1 {
		if y&1 == 1 {
			r = r.mul(x)
		}
	}
	return r
}
