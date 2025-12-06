package math

import . "codeforces-go/common"

const M = 998244353

func Norm(x int) int {
	x %= M
	if x < 0 {
		x += M
	}
	return x
}

func Add(x, y int) int {
	r := x + y
	if r >= M {
		r -= M
	}
	return r
}

func Sub(x, y int) int {
	r := x - y
	if r < 0 {
		r += M
	}
	return r
}

func Mul(x, y int) int {
	return Norm(x * y)
}

func Inv(x int) int {
	return Pow(x, M-2, M)
}

func Div(x, y int) int {
	return Mul(x, Inv(y))
}
