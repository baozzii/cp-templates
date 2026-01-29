package templates

func abs[T real](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}

func gcd[T integer](x, y T) T {
	if x < 0 || y < 0 {
		return gcd(abs(x), abs(y))
	}
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func lcm[T integer](x, y T) T {
	return x / gcd(x, y) * y
}

func pow[S, T integer](x S, n T, m S) S {
	r := S(1)
	for ; n > 0; n, x = n>>1, x*x%m {
		if n%2 == 1 {
			r = r * x % m
		}
	}
	return r
}

func exgcd[T integer](a, b T) (T, T, T) {
	if b == 0 {
		return a, 1, 0
	}
	d, x2, y2 := exgcd(b, a%b)
	return d, y2, x2 - (a/b)*y2
}
