package trees

type LCA struct {
	t  *Tree
	st [][]int32
	lg []int32
}

func NewLCA(t *Tree) *LCA {
	n := t.n
	lg := make([]int32, n+1)
	for i := int32(2); i <= n; i++ {
		lg[i] = lg[i>>1] + 1
	}
	K := lg[n]
	st := make([][]int32, K+1)
	st[0] = make([]int32, n)
	for pos := int32(0); pos < n; pos++ {
		u := t.ord[pos]
		st[0][pos] = int32(t.pa[u])
	}
	get := func(x, y int32) int32 {
		if x == -1 {
			return y
		}
		if y == -1 {
			return x
		}
		if t.in[x] < t.in[y] {
			return x
		}
		return y
	}
	for k := int32(1); k <= K; k++ {
		length := int32(1) << k
		half := length >> 1
		st[k] = make([]int32, n-length+1)
		for i := int32(0); i+length <= n; i++ {
			st[k][i] = get(st[k-1][i], st[k-1][i+half])
		}
	}
	return &LCA{t: t, st: st, lg: lg}
}

func (s *LCA) lca(u, v int) int {
	if u == v {
		return u
	}
	t := s.t
	du, dv := t.in[u], t.in[v]
	if du > dv {
		du, dv = dv, du
	}
	L := du + 1
	R := dv
	length := R - L + 1
	k := s.lg[length]
	row := s.st[k]
	x := row[L]
	y := row[R-(1<<k)+1]
	if x == -1 {
		return int(y)
	}
	if y == -1 {
		return int(x)
	}
	if t.in[x] < t.in[y] {
		return int(x)
	}
	return int(y)
}
