package templates

import "math/bits"

func ctz[T integer](x T) int {
	return bits.TrailingZeros(uint(x))
}

func clz[T integer](x T) int {
	return bits.LeadingZeros(uint(x))
}

func popcount[T integer](x T) int {
	return bits.OnesCount(uint(x))
}

func lowbit[T integer](x T) T {
	y := uint(x)
	return T(y & -y)
}

func highbit[T integer](x T) T {
	if x == 0 {
		return x
	}
	return T(1) << (63 - clz(uint(x)))
}
