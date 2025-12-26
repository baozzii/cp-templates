package strings

import "cmp"

func MinExpression[T cmp.Ordered, E ~[]T](a E) int {
	k, i, j := 0, 0, 1
	n := len(a)
	for k < n && i < n && j < n {
		if a[(i+k)%n] == a[(j+k)%n] {
			k++
		} else {
			if a[(i+k)%n] > a[(j+k)%n] {
				i += k + 1
			} else {
				j += k + 1
			}
			if i == j {
				i++
			}
			k = 0
		}
	}
	return min(i, j)
}

func MinExpressionFunc[T any, E ~[]T](a E, cmp func(T, T) int) int {
	k, i, j := 0, 0, 1
	n := len(a)
	for k < n && i < n && j < n {
		c := cmp(a[(i+k)%n], a[(j+k)%n])
		if c == 0 {
			k++
		} else {
			if c > 0 {
				i += k + 1
			} else {
				j += k + 1
			}
			if i == j {
				i++
			}
			k = 0
		}
	}
	return min(i, j)
}
