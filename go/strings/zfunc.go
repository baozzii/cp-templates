package strings

func ZFunc[T comparable, E ~[]T](a E) []int {
	n := len(a)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r && z[i-l] < r-i+1 {
			z[i] = z[i-l]
		} else {
			z[i] = max(0, r-i+1)
			for i+z[i] < n && a[z[i]] == a[i+z[i]] {
				z[i]++
			}
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

func ZFuncFunc[T any, E ~[]T](a E, eq func(T, T) bool) []int {
	n := len(a)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r && z[i-l] < r-i+1 {
			z[i] = z[i-l]
		} else {
			z[i] = max(0, r-i+1)
			for i+z[i] < n && eq(a[z[i]], a[i+z[i]]) {
				z[i]++
			}
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}
