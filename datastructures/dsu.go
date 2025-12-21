package datastructures

type DSU struct {
	fa, sz []int
}

func NewDSU(n int) *DSU {
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	return &DSU{fa, sz}
}

func (dsu *DSU) Find(u int) int {
	if u == dsu.fa[u] {
		return u
	}
	dsu.fa[u] = dsu.Find(dsu.fa[u])
	return dsu.fa[u]
}

func (dsu *DSU) Merge(u, v int) bool {
	u = dsu.Find(u)
	v = dsu.Find(v)
	if u == v {
		return false
	}
	if dsu.sz[u] > dsu.sz[v] {
		u, v = v, u
	}
	dsu.sz[v] += dsu.sz[u]
	dsu.fa[u] = v
	return true
}

func (dsu *DSU) Same(u, v int) bool {
	return dsu.Find(u) == dsu.Find(v)
}

func (dsu *DSU) Groups() [][]int {
	g := make([][]int, len(dsu.fa))
	for i := range g {
		g[dsu.Find(i)] = append(g[dsu.Find(i)], i)
	}
	r := make([][]int, 0)
	for i := range g {
		if len(g[i]) > 0 {
			r = append(r, g[i])
		}
	}
	return r
}
