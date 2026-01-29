package templates

type dsu struct {
	fa, sz []int32
}

func new_dsu(n int) *dsu {
	fa := make([]int32, n)
	sz := make([]int32, n)
	for i := range fa {
		fa[i] = int32(i)
		sz[i] = 1
	}
	return &dsu{fa, sz}
}

func (dsu *dsu) find(_u int) int {
	u := int(_u)
	if u == int(dsu.fa[u]) {
		return u
	}
	dsu.fa[u] = int32(dsu.find(int(dsu.fa[u])))
	return int(dsu.fa[u])
}

func (dsu *dsu) merge(u, v int) bool {
	u = dsu.find(u)
	v = dsu.find(v)
	if u == v {
		return false
	}
	if dsu.sz[u] > dsu.sz[v] {
		u, v = v, u
	}
	dsu.sz[v] += dsu.sz[u]
	dsu.fa[u] = int32(v)
	return true
}

func (dsu *dsu) same(u, v int) bool {
	return dsu.find(u) == dsu.find(v)
}

func (dsu *dsu) groups() [][]int {
	g := make([][]int, len(dsu.fa))
	for i := range g {
		g[dsu.find(i)] = append(g[dsu.find(i)], i)
	}
	r := make([][]int, 0)
	for i := range g {
		if len(g[i]) > 0 {
			r = append(r, g[i])
		}
	}
	return r
}
