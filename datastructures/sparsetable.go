package datastructures

/*
https://judge.yosupo.jp/submission/334191
*/

import . "cp-templates/common"

type SparseTableInfo[T any] interface {
	op(T, T) T
	e() T
}

type SparseTable[T any, M SparseTableInfo[T]] struct {
	t  [][]T
	op func(T, T) T
	e  func() T
}

func NewSparseTableWith[T any, M SparseTableInfo[T]](a []T, m M) *SparseTable[T, M] {
	op := m.op
	e := m.e
	n := len(a)
	l := 63 - Clz(n)

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
	return &SparseTable[T, M]{st, op, e}
}

func NewSparseTable[T any, M SparseTableInfo[T]](n int, m M) *SparseTable[T, M] {
	a := make([]T, n)
	for i := range a {
		a[i] = m.e()
	}
	return NewSparseTableWith(a, m)
}

func (st *SparseTable[T, M]) Query(l, r int) T {
	if l == r {
		return st.e()
	}
	i := 63 - Clz(r-l)
	return st.op(st.t[i][l], st.t[i][r-(1<<i)])
}
