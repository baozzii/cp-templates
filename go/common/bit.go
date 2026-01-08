package common

import "math/bits"

func Ctz[T Integer](x T) int {
	return bits.TrailingZeros(uint(x))
}

func Clz[T Integer](x T) int {
	return bits.LeadingZeros(uint(x))
}

func Popcount[T Integer](x T) int {
	return bits.OnesCount(uint(x))
}

func Lowbit[T Integer](x T) T {
	y := uint(x)
	return T(y & -y)
}

func Highbit[T Integer](x T) T {
	if x == 0 {
		return x
	}
	return T(1) << (63 - Clz(uint(x)))
}
