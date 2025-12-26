package strings

func GetPi[T comparable, E ~[]T](s E) []int {
	n := len(s)
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

func GetPiFunc[T any, E ~[]T](s E, eq func(T, T) bool) []int {
	n := len(s)
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && !eq(s[i], s[j]) {
			j = pi[j-1]
		}
		if eq(s[i], s[j]) {
			j++
		}
		pi[i] = j
	}
	return pi
}
