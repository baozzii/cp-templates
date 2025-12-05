package common

func Abs[T RealNumber](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}

func Gcd[T Integer](x, y T) T {
	if x < 0 || y < 0 {
		return Gcd(Abs(x), Abs(y))
	}
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

func Lcm[T Integer](x, y T) T {
	return x / Gcd(x, y) * y
}

func Pow[S, T Integer](x S, n T, m S) S {
	r := S(1)
	for ; n > 0; n, x = n>>1, x*x%m {
		if n%2 == 1 {
			r = r * x % m
		}
	}
	return r
}
