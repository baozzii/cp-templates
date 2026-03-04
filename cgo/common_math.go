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
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func lcm[T integer](x, y T) T {
	return x / gcd(x, y) * y
}

func pow[S, T integer](x S, n T, m S) S {
	r := S(1) % m
	for ; n > 0; n, x = n>>1, x*x%m {
		if n%2 == 1 {
			r = r * x % m
		}
	}
	return r
}

func exgcd[T integer](a, b T) (T, T, T) {
	x0, y0 := T(1), T(0)
	x1, y1 := T(0), T(1)
	for b != 0 {
		q := a / b
		a, b = b, a-q*b
		x0, x1 = x1, x0-q*x1
		y0, y1 = y1, y0-q*y1
	}
	return a, x0, y0
}
