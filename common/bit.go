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
