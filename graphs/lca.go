package graphs

import . "codeforces-go/common"

type LCA struct {
	t  *Tree
	st [][]int
}

func NewLCA(t *Tree) *LCA {
	n := t.n
	l := 63 - Clz(n)
	st := make([][]int, l+1)
	for i := range st {
		st[i] = make([]int, n)
		for j := range st[i] {
			st[i][j] = -1
		}
	}
	for i := 0; i < n; i++ {
		st[0][t.in[i]] = i
	}
	for j := 1; j <= l; j++ {
		for i := 0; i+(1<<j) <= n; i++ {
			u, v := st[j-1][i], st[j-1][i+(1<<(j-1))]
			var w int
			if t.dep[u] > t.dep[v] {
				w = v
			} else {
				w = u
			}
			st[j][i] = w
		}
	}
	return &LCA{t, st}
}

func (s *LCA) lca(u, v int) int {
	if u == v {
		return u
	}
	u = s.t.in[u]
	v = s.t.in[v]
	if u > v {
		u, v = v, u
	}
	u++
	v++
	i := 63 - Clz(v-u)
	var w int
	if s.t.dep[s.st[i][u]] > s.t.dep[s.st[i][v-(1<<i)]] {
		w = s.st[i][v-(1<<i)]
	} else {
		w = s.st[i][u]
	}
	return s.t.pa[w]
}
