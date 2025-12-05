package datastructures

import "math/bits"

type sparsetable_info[T any] interface {
	op(T, T) T
	e() T
}

type sparsetable[T any, M sparsetable_info[T]] struct {
	t  [][]T
	op func(T, T) T
	e  func() T
}

func new_sparsetable_with[T any, M sparsetable_info[T]](a []T, m M) *sparsetable[T, M] {
	op := m.op
	e := m.e
	n := len(a)
	l := 63 - bits.LeadingZeros(uint(n))

	st := make([][]T, l+1)
	for i := range st {
		st[i] = make([]T, n)
		for j := range st[i] {
			st[i][j] = e()
		}
	}
	copy(st[0], a)
	for j := 1; j <= l; j++ {
		for i := 0; i+(1<<j) <= n; i++ {
			st[j][i] = op(st[j-1][i], st[j-1][i+(1<<(j-1))])
		}
	}
	return &sparsetable[T, M]{st, op, e}
}

func new_sparsetable[T any, M sparsetable_info[T]](n int, m M) *sparsetable[T, M] {
	a := make([]T, n)
	for i := range a {
		a[i] = m.e()
	}
	return new_sparsetable_with(a, m)
}

func (st *sparsetable[T, M]) query(l, r int) T {
	if l == r {
		return st.e()
	}
	i := 63 - bits.LeadingZeros(uint(r-l))
	return st.op(st.t[i][l], st.t[i][r-(1<<i)])
}
