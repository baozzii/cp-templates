package strings

type Manacher struct {
	d1, d2 []int32
}

func (m *Manacher) IsPal(l, r int) bool {
	if l == r || l+1 == r {
		return true
	}
	d := r - l
	if d%2 == 1 {
		return int(m.d1[(l+r-1)/2]) >= d/2+1
	} else {
		return int(m.d2[(l+r)/2]) >= d/2
	}
}

func (m *Manacher) MaxLen(i, c int) int {
	if c == 1 {
		return int(m.d1[i])*2 - 1
	} else {
		return int(m.d2[i]) * 2
	}
}

func NewManacher[T comparable, E ~[]T](s E) *Manacher {
	n := len(s)
	d1 := make([]int32, n)
	for i, l, r := 0, 0, -1; i < n; i++ {
		var k int
		if i > r {
			k = 1
		} else {
			k = min(int(d1[l+r-i]), r-i+1)
		}
		for i-k >= 0 && i+k < n && s[i-k] == s[i+k] {
			k++
		}
		d1[i] = int32(k)
		k--
		if i+k > r {
			l = i - k
			r = i + k
		}
	}
	d2 := make([]int32, n)
	for i, l, r := 0, 0, -1; i < n; i++ {
		var k int
		if i > r {
			k = 0
		} else {
			k = min(int(d2[l+r-i+1]), r-i+1)
		}
		for i-k-1 >= 0 && i+k < n && s[i-k-1] == s[i+k] {
			k++
		}
		d2[i] = int32(k)
		k--
		if i+k > r {
			l = i - k - 1
			r = i + k
		}
	}
	return &Manacher{d1, d2}
}

func NewManacherFunc[T any, E ~[]T](s E, eq func(T, T) bool) *Manacher {
	n := len(s)
	d1 := make([]int32, n)
	for i, l, r := 0, 0, -1; i < n; i++ {
		var k int
		if i > r {
			k = 1
		} else {
			k = min(int(d1[l+r-i]), r-i+1)
		}
		for i-k >= 0 && i+k < n && eq(s[i-k], s[i+k]) {
			k++
		}
		d1[i] = int32(k)
		k--
		if i+k > r {
			l = i - k
			r = i + k
		}
	}
	d2 := make([]int32, n)
	for i, l, r := 0, 0, -1; i < n; i++ {
		var k int
		if i > r {
			k = 0
		} else {
			k = min(int(d2[l+r-i+1]), r-i+1)
		}
		for i-k-1 >= 0 && i+k < n && eq(s[i-k-1], s[i+k]) {
			k++
		}
		d2[i] = int32(k)
		k--
		if i+k > r {
			l = i - k - 1
			r = i + k
		}
	}
	return &Manacher{d1, d2}
}
